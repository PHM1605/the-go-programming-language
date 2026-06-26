package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// -------- NEW: Cancellation implementation --------
// cancellation channel: nothing will be sent here; only to "close()"
var done = make(chan struct{})

// polling to see if channel "done" is closed or not
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// return the entries of directory "dir"
// NOTE: to limit #goroutines that call "dirents" is called => we wrap it around a SEMAPHORE
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.DirEntry {
	// NEW: check Cancellation in between
	select {
	case sema <- struct{}{}: // put token to 20-slot semaphore
		// do nothing else
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release that token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// scan entries of a dir, inject sizes OF FILES (not dir) into a channel
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	// reduct goroutine count when done
	defer n.Done()
	// check if User is done with this program
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			// NEW
			n.Add(1)

			subdir := filepath.Join(dir, entry.Name())
			// NEW: extra goroutine
			go walkDir(subdir, n, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
				continue
			}
			fileSizes <- info.Size()
		}
	}
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// get command line directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// channel to pump file size
	fileSizes := make(chan int64)
	// NEW: try concurrent with "n" as goroutine counter
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	// wait to close after all goroutines finished
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// NEW: when user press something => close "done" channel => cancelled() is "false"
	go func() {
		os.Stdin.Read(make([]byte, 1)) // create a byte slice of 1 element
		close(done)                    // trigger "case <-done" in our main for-loop
	}()

	// channel to send events every 500ms
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// handle 2 channels with "select"
	// run it to check STATUS frequently
	var nfiles, nbytes int64

loop: // NOTE: syntax to "break" for loop wrapping around a "select"
	for {
		select {
		// NEW: if User press something to cancel all process
		case <-done:
			// Drain fileSizes to to allow everything finish
			for range fileSizes {
				// Do nothing
			}
			return

		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

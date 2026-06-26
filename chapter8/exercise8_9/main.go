package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// file size of a file AND its root (which root from command line)
// e.g. fileA has size 2GB and belongs to root /usr; folderB has size 3GB and belongs to root /home
// => push {2,"/usr"} and {3,"/home"} to channel
type fileSize struct {
	root string
	size int64
}

// semaphores limits #goroutines calling this function
var sema = make(chan struct{}, 20)

// return entries in this folder
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // place token on 20-slot semaphore
	defer func() { <-sema }() // release token for empty space

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}

// NEW: now when we walk through a Directory, we must know what is its ROOT
// waitGroup: goroutine counter
// fileSizes: receiving
func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- fileSize) {
	// make done for this goroutine of walkDir() => counter minus 1
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			// start a new goroutine for the subdir
			n.Add(1)
			go walkDir(root, subdir, n, fileSizes)
		} else {
			// actual file size calculation for FILE
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du: %v\n", err)
				continue
			}
			// inject file size AND its ROOT to channel
			fileSizes <- fileSize{
				root: root,
				size: info.Size(),
			}
		}
	}
}

// flag to print status every 500ms
var verbose = flag.Bool("v", false, "show verbose progress messages")

// print values of 2 maps: rootA with #files and #bytes at that time of gathering
func printDiskUsage(nfiles map[string]int64, nbytes map[string]int64) {
	for root := range nfiles {
		fmt.Printf(
			"%s: %d files %.1f GB\n",
			root,
			nfiles[root],
			float64(nbytes[root])/1e9,
		)
	}
	fmt.Println()
}

func main() {
	// get list of root directories from command line
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan fileSize) // file size and it's root name

	// goroutine counter to properly close them at the end
	var n sync.WaitGroup
	// get filesize (with root name) info
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	// properly close goroutines
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// sending channel to tick "status printing" every 500ms
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// root directory "A" has how many files
	nfiles := make(map[string]int64)
	// root directory "A" has size of how many bytes
	nbytes := make(map[string]int64)

	// loop: choose where to break (break loop)
loop:
	// "scanning status and file size (with root name) gathering" for loop
	for {
		select {
		// file size (with root name) gathering
		case fs, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles[fs.root]++
			nbytes[fs.root] += fs.size
		// scanning status
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

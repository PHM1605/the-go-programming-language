package main

import (
	"context"
	"exercise8_10/links"
	"fmt"
	"log"
	"os"
	"time"
)

// crawl a link for links of AND inside it
// NEW: use "context" for cancellation a goroutine
func crawl(ctx context.Context, url string) []string {
	// we simply print the link we found
	fmt.Println(url)
	list, err := links.Extract(ctx, url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// channel that process all links => de-duplicate
	worklist := make(chan []string)
	// send each non-duplicated link one-by-one
	unseenLinks := make(chan string)

	// NEW: cancellation context
	// ctx.Done() will create a "cancellation channel"
	// every time we call "cancel()" that cancellation channel will ignite "0" i.e. "case <-ctx.Done()" will run
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// push to raw worklist channel what from command line first
	go func() {
		worklist <- os.Args[1:]
	}()

	// NEW: try triggering "cancel" after 5s
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("canceling...")
		cancel()
	}()

	// create 20 crawlers only
	for i := 0; i < 20; i++ {
		go func() {
			// crawl "link" that pops here multiple times
			for link := range unseenLinks {
				foundLinks := crawl(ctx, link)
				// update list of raw links OR cancel
				select {
				case worklist <- foundLinks:
					// do nothing
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// set of non-duplicated links
	seen := make(map[string]bool)
	// to keep track of 2 channels => use "select"
	// to keep receiving from "worklist" => use outer "for" loop
	for {
		select {
		case <-ctx.Done():
			fmt.Println("crawl canceled")
			return
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					// additional "select" to cancel hanging here too
					select {
					case unseenLinks <- link:
						// do nothing
					case <-ctx.Done():
						return
					}
				}
			}
		}

	}
}

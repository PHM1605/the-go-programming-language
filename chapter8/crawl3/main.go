package main

import (
	"crawl3/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	// we only PRINT the list of URLs in the crawling process
	fmt.Println(url)
	// list of links in that URL (and its children)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// create a channel to push list of links to crawl (may have duplicates)
	worklist := make(chan []string)
	// de-duplicates URLs
	unseenLinks := make(chan string)

	// push to that channel what from command line first
	go func() { worklist <- os.Args[1:] }()

	// NEW: create 20 goroutines as crawlers only
	for i := 0; i < 20; i++ {
		go func() {
			// loop over each link (1 link each time, only stop when "unseenLinks" is no longer pushed)
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// NOTE: we must create a new goroutine to modify "worklist"
				// otherwise it blocks here to wait for reading from "worklist" forever
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// de-duplicate worklist items
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link // jump to the PREVIOUS paragraph "range unseenLinks"
			}
		}
	}
}

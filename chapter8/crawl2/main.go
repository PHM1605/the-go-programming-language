package main

import (
	"crawl1/links"
	"fmt"
	"log"
	"os"
)

// NEW: "tokens channel": is a counting semaphore; 20 slots to limit 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	// we only PRINT the list of URLs in the crawling process
	fmt.Println(url)

	// NEW: say that "hey, I take that slot"
	tokens <- struct{}{}

	// list of links in that URL (and its children)
	list, err := links.Extract(url)

	// NEW: empty that slot to show that other process can enter
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// create a channel to push list of links to crawl (set-links-1=>goroutine-1, set-links-2=>goroutine-2 etc.)
	worklist := make(chan []string)
	// NEW: number of pending sends to worklist
	var n int

	// push to that channel what from command line first
	// NOTE: we must create goroutine here; otherwise it blocks when nothing is received
	n++
	go func() { worklist <- os.Args[1:] }()

	// set of links to crawl
	seen := make(map[string]bool)

	// NEW: as soon as all pushed goroutines finished => this loop escapes
	for ; n > 0; n-- {
		// proceed each set of links CONCURRENTLY
		list := <-worklist // (take 1 set out of channel)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link) // push result back to worklist to crawl deeper
				}(link)
			}
		}
	}
}

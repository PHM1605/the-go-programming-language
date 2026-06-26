package main

import (
	"crawl1/links"
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
	// create a channel to push list of links to crawl (set-links-1=>goroutine-1, set-links-2=>goroutine-2 etc.)
	worklist := make(chan []string)

	// push to that channel what from command line first
	// NOTE: we must create goroutine here; otherwise it blocks when nothing is received
	go func() { worklist <- os.Args[1:] }()

	// set of links to crawl
	seen := make(map[string]bool)
	for list := range worklist {
		// proceed each set of links CONCURRENTLY
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link) // push result back to worklist to crawl deeper
				}(link)
			}
		}
	}
}

package main

import (
	"exercise8_6/links"
	"flag"
	"fmt"
	"log"
)

// for flag -depth=3
var maxDepth = flag.Int("depth", 3, "maximum crawl depth")

// NEW: we add also the depth of each item
type item struct {
	url   string
	depth int
}

func crawl(url string) []string {
	// for this crawling we only PRINT
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()

	worklist := make(chan []item)
	unseenLinks := make(chan item) // after deduplicate, eventual link will be sent here

	// first, send command-line arguments
	go func() {
		var items []item
		for _, url := range flag.Args() {
			items = append(items, item{url, 0})
		}
		worklist <- items
	}()

	// 10 web crawlers
	for i := 0; i < 10; i++ {
		go func() {
			// take ONE item each time for each crawler
			for it := range unseenLinks {
				// don't crawl if depth > 3
				if it.depth >= *maxDepth {
					continue
				}
				foundLinks := crawl(it.url)
				// fix links => item format
				var items []item
				for _, link := range foundLinks {
					items = append(items, item{
						url:   link,
						depth: it.depth + 1,
					})
				}
				// NOTE: every time we change worklist we must create a new goroutine
				go func() {
					worklist <- items
				}()
			}
		}()
	}

	// deduplicate links before sending to "unseenLinks" channel, one by one
	seen := make(map[string]bool)
	for list := range worklist {
		for _, it := range list {
			if !seen[it.url] {
				seen[it.url] = true
				unseenLinks <- it
			}
		}
	}
}

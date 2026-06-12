package main

import (
	"findlinks3/links"
	"fmt"
	"log"
	"os"
)

// 1st parameter: a function takes a "string" and returns a "[]string" => crawl LINKS from an URL
// 2nd parameter: []string; list of websites
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // make which websites have been done

	// 1st iteration: crawl to get a list of URLs from "https://golang.org" and "https://google.com" => [https://golang.org/users, https://google.com/templates]
	// 2nd iteration: dig deeper => [https://golang.org/users/1, https://golang.org/users/abc, https://google.com/templates/xyz]
	// 3rd iteration: crawl(item) finds nothing => escape for loop
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		// iterate over list of websites
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawling function: takes an URL an return list of links
func crawl(url string) []string {
	// we only PRINT in this exercise
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list // list of links from that URL
}

func main() {
	// we can "crawl" of a list of websites ["https://golang.org", "https://google.com"]
	breadthFirst(crawl, os.Args[1:])
}

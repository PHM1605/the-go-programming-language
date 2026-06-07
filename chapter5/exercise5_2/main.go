package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func countElements(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	countElements(counts, n.FirstChild)
	countElements(counts, n.NextSibling)
}

func main() {
	file, err := os.Open("exercise5_2.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}
	// number of <div>, <span> etc.
	counts := make(map[string]int)
	countElements(counts, doc)
	// print result
	for tag, count := range counts {
		fmt.Printf("%-10s %d\n", tag, count)
	}
}

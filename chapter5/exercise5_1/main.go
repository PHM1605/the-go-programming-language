package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// traverse HTML node tree, extract link from <a href='...'>, append to []string
func visit(links []string, n *html.Node) []string {
	// NEW: end of recursion
	if n == nil {
		return links
	}
	// if current Node is "anchor"
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// NEW: traverse BOTH Child and Siblings
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// traverse HTML Node tree
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

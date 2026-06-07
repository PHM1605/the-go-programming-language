package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// traverse HTML node tree, extract link from <a href='...'>, append to []string
func visit(links []string, n *html.Node) []string {
	// if current Node is "anchor"
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
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

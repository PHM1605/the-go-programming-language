package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// apply a <pre> and <post> function for each Node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// MIDDLE: what is actually done on this Node
	// in this case, we do nothing but only recursion (all printing has been done BEFORE and AFTER with PRE and POST functions)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func outline(doc *html.Node) {
	// NEW: depth is now a shared variable of 2 ANONYMOUS FUNCTIONs
	depth := 0 // how deep is our indentation for each line

	// 2 anonymous functions
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	// print <pre> -> dig deep -> print <pos>
	forEachNode(doc, startElement, endElement)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
	}
	// GET request to website
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "get: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// JSON to Node
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}
	// give outline to a Node and it's children
	outline(doc)
}

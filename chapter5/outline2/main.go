package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int // indent from beginning of each HTML element

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

// print starting tag <div> for each Node
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// "%*s, 6" == "%6s"; then print that 6-space with a string ""
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

// print ending tag </div> for each Node
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
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
	// perform a "PRE-FUNCTION" and "POST-FUNCTION" for that Node
	forEachNode(doc, startElement, endElement)
}

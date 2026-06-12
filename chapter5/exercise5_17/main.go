package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// names: list of tags we wanna search
func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var result []*html.Node
	// set of names we are looking for
	wanted := make(map[string]bool)
	for _, name := range names {
		wanted[name] = true
	}
	// recursive visit
	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n == nil {
			return
		}
		// actual work
		if n.Type == html.ElementNode && wanted[n.Data] {
			result = append(result, n)
		}
		// visit children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return result
}

func main() {
	// open html file
	f, err := os.Open("exercise5_17.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// convert html file to tree
	doc, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	// try finding data
	fmt.Println("Images:")
	images := ElementsByTagName(doc, "img")
	for _, n := range images {
		fmt.Println(n.Data)
	}

	// try finding headings
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println("\nHeadings:")
	for _, n := range headings {
		fmt.Println(n.Data)
	}
}

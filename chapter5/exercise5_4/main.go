package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	// current node processing
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img", "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	// recursion
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {
	file, err := os.Open("exercise5_4.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

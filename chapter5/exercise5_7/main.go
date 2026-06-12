package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int // indent from beginning of each HTML element

func prettyPrint(n *html.Node) {
	if n == nil {
		return
	}
	// for current Node
	switch n.Type {
	case html.ElementNode:
		// self enclosing elements
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=%q", a.Key, a.Val)
			}
			fmt.Println("/>")
			return
		}

		// Elements with <div> and </div>
		// opening tag
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		fmt.Println(">")
		// what inside of tag
		depth++
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			prettyPrint(c)
		}
		depth--
		// closing tag
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}

	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)

	default:
		// print over the children nodes RECURSIVE
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			prettyPrint(c)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
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

	prettyPrint(doc)
}

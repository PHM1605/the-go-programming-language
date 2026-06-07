package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func printText(n *html.Node) {
	if n == nil {
		return
	}
	// skip "script" and "style" node
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		printText(n.NextSibling) // NOTE: those nodes don't have "Child" but have "Sibling"
		return
	}
	// print TextNode
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Println(text)
		}
	}
	// Recursive
	printText(n.FirstChild)
	printText(n.NextSibling)
}

func main() {
	file, err := os.Open("exercise5_3.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	printText(doc)
}

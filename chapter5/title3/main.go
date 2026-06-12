package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Walk through HTML tree
func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	// NOTE: we define here a <type> for "multiple-header-panic"
	type bailout struct{}

	// NOTE: we try to "recover" "panic" in lower-part here
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// expected panic => convert to "error"
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // not the expected type => we "throw panic" again
		}
	}()

	// our usual walking HTML code; provide only <pre> function, <post> nil
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			// NOTE: at 1st <title> this "title" variable is always ""
			if title != "" {
				panic(bailout{}) // NOTE: create a panic that we have expected
			}
			// we make a trick at the end: change value of that "title" var to non-nil
			title = n.FirstChild.Data
		}
	}, nil)
	// <header> with no <title> is also error
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	// here we test a HTML with TWO <title>
	htmlText := `
		<html>
		<head>
			<title>First Title</title>
			<title>Second Title</title>
		<head>
		<body>
			<h1>Hello</h1>
		</body>
	`
	doc, err := html.Parse(strings.NewReader(htmlText))
	if err != nil {
		panic(err)
	}
	title, err := soleTitle(doc)
	fmt.Println("title", title)
	fmt.Println("error:", err)
}

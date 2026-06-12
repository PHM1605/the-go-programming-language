package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int // indent from beginning of each HTML element

// apply a <pre> and <post> function for each Node
// NEW: bool=true => continue traversal; false => stop immediately
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if n == nil {
		return true
	}
	// searching <id> here; if found then not propagate => return "false" immediately
	if pre != nil && !pre(n) {
		return false
	}

	// MIDDLE: what is actually done on this Node
	// in this case, we do nothing but only recursion (all printing has been done BEFORE and AFTER with PRE and POST functions)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// stop progate if EXECUTE and SEE FALSE flag
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	// doing nothing here
	if post != nil && !post(n) {
		return false
	}
	// propagate deeper
	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node
	// function before each Node; if "false" then not propagate deeper
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		// check if this ElementNode is the "id" we are searching
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				result = n   // NOTE: we FOUND here
				return false // no propagate anymore
			}
		}
		return true // auto-propagate
	}
	// do the "search-and-propagate stuff"
	forEachNode(doc, pre, nil)
	return result
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
	file, err := os.Open("exercise5_8.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	// Look for item named "footer"
	node := ElementByID(doc, "footer")
	if node != nil {
		fmt.Printf("Found <%s>\n", node.Data)
	} else {
		fmt.Println("Not found")
	}
}

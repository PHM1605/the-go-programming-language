package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// get elements in the Tree
type Element struct {
	Name  string
	Id    string
	Class string
}

// get command line arguments
type Selector struct {
	Name  string
	Id    string
	Class string
}

// parse arguments like "div", "div#page", "div.wide", "div#page.wide"
func parseSelector(s string) Selector {
	var sel Selector
	// class
	if strings.Contains(s, ".") {
		parts := strings.SplitN(s, ".", 2) // at most 2 substrings
		s = parts[0]                       // still "div#page"
		sel.Class = parts[1]
	}
	// get id
	if strings.Contains(s, "#") { // still "div#page"
		parts := strings.SplitN(s, "#", 2) // at most 2 substrings
		sel.Name = parts[0]
		sel.Id = parts[1]
	} else {
		sel.Name = s
	}
	return sel
}

// attrs: list of [{Name:"id",Value:"page"}, {Name:"class",Value:"wide"}, ...]
// get value of that attribute named e.g. "class"
func getAttr(attrs []xml.Attr, name string) string {
	for _, a := range attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

// check if e.g. "#page.wide" matches "div.wide" (already parsed into Element and Selector)
func matches(e Element, s Selector) bool {
	// empty Name in the argument means "not consider"
	if s.Name != "" && e.Name != s.Name {
		return false
	}
	if s.Id != "" && e.Id != s.Id {
		return false
	}
	if s.Class != "" && e.Class != s.Class {
		return false
	}
	return true
}

// check if stack <x> contains all selectors in arguments y
func containsAll(x []Element, y []Selector) bool {
	// if arguments to check is shorter than stack
	for len(y) <= len(x) {
		// if we check all arguments
		if len(y) == 0 {
			return true
		}
		if matches(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:] // shorten the stack to check deeper
	}
	return false
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element      // NEW: we stock all attributes, not only string name
	var selectors []Selector // from command line
	for _, arg := range os.Args[1:] {
		selectors = append(selectors, parseSelector(arg))
	}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, Element{
				Name:  tok.Name.Local,
				Id:    getAttr(tok.Attr, "id"),
				Class: getAttr(tok.Attr, "class"),
			})
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop that element out
		case xml.CharData:
			text := strings.TrimSpace(string(tok))
			if text == "" {
				continue
			}

			if containsAll(stack, selectors) {
				// combine stack elements only for printing out
				var names []string
				for _, e := range stack {
					names = append(names, e.Name)
				}
				fmt.Printf("%s: %s\n", strings.Join(names, " "), text)
			}
		}

	}
}

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// Node Interface has 2 types: CharData or Element
type Node interface{}
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr // id="users" class="registerInfo"
	Children []Node
}

func printTree(n Node, depth int) {
	indent := strings.Repeat("--", depth)
	switch n := n.(type) {
	case CharData:
		text := strings.TrimSpace(string(n))
		if text != "" {
			fmt.Printf("%s%q\n", indent, text)
		}
	case *Element:
		fmt.Printf("%s<%s", indent, n.Type.Local) // <div
		for _, attr := range n.Attr {
			fmt.Printf(" %s=%q", attr.Name.Local, attr.Value) // class="navBar" style="xyz"
		}
		fmt.Printf(">\n") // >
		for _, child := range n.Children {
			printTree(child, depth+1)
		}
	}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element // list of PARENT and GRANDPARENT etc.
	var root Node

	for {
		tok, err := dec.Token() // tok: XML Token Interface
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "xmltree: %v\n")
			os.Exit(1)
		}

		switch tok := tok.(type) {

		case xml.StartElement:
			el := &Element{
				Type: tok.Name, // "div"
				Attr: tok.Attr, // // id="users" class="registerInfo"
			}
			// if stack empty, this is root
			if len(stack) == 0 {
				root = el
			} else {
				// push new Element to last Parent (in Stack)'s list of Children
				// e.g. <div1>
				//				<div2>
				//					<h2></h2>
				//					<added div> here, then Stack is only TWO <div> on top (<h2> has been popped out)
				parent := stack[len(stack)-1] // <div2>
				parent.Children = append(parent.Children, el)
			}
			// push new Element in Stack
			stack = append(stack, el)

		case xml.EndElement:
			// pop that element out
			stack = stack[:len(stack)-1]

		case xml.CharData:
			// Create a Text Element named "char"
			text := strings.TrimSpace(string(tok))
			if text == "" {
				continue
			}
			char := CharData(text)

			if len(stack) > 0 {
				// last parent in Stack
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, char)
			}
		}
	}
	printTree(root, 0) // depth of <root> is 0 (to print indent)
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"encoding/xml"
)

// check if x (stack) contains elements of y (what in arguments)
// e.g. x: ["div","div","h2","div"]; y: ["div","div","h2"]
func containsAll(x, y []string) bool {
	// stack x must always be > requirements in arguments y
	for len(y) <= len(x) {
		// we check all arguments y
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
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
			stack = append(stack, tok.Name.Local) // push that element e.g. <div> to Stack
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop that Element </div> out
		case xml.CharData: // content inside that <div>
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

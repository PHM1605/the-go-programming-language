// like "basename()" function => a.go is "a", a/b_c.go is "b_c"
// go run basename2.go
package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	// remove last "/" and what before
	s = s[slash+1:]
	// remove .go extension
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	tests := []string{
		"a.go",
		"a/b_c.go",
		"/usr/local/bin/test.txt",
		"no_extension",
		"multi.part.name.go",
	}
	for _, t := range tests {
		fmt.Printf("basename(%q) = %q\n", t, basename(t))
	}
}

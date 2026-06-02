// like "basename()" function => a.go is "a", a/b_c.go is "b_c"
// go run basename1.go
package main

import "fmt"

func basename(s string) string {
	// Discard "/" and before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve anything before .go extension
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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

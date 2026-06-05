// Write in-place function to SQUASH adjacent Unicode spaces (' ', '\t' etc.) in []byte slice. =>
// go run exercise4_6.go
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	// i will scan through the slice ([]byte)
	// w: write index
	w := 0
	inSpace := false // when <i> is in middle of a long space => do not copy
	for i := 0; i < len(b); {
		// decode each rune
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			// only add 1st space i.e. when <inSpace> is still "false"
			if !inSpace {
				b[w] = ' '
				w++
				inSpace = true
			}
		} else {
			// copy the rune's bytes to BACKWARD
			copy(b[w:], b[i:i+size])
			w += size
			inSpace = false
		}
		i += size
	}
	return b[:w]
}

func main() {
	s := []byte("Hello\t\tworld \n\n Go\u00A0lang")
	fmt.Println(string(squashSpaces(s))) // "Hello world Go lang"
}

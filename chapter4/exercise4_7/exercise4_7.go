// Write in-place function to reverse UTF-8 string (so each char can be more than 1 byte)
// go run exercise4_7.go
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUTF8(b []byte) {
	// reverse each rune individually 世 = E4 B8 96 => 96 B8 E4
	// so that later when we reverse the WHOLE <b> rune pattern switches back to E4 B8 96
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}
	// reverse entire slice
	reverseBytes(b)
}

func main() {
	s := []byte("Hello, 世界")
	reverseUTF8(s)
	fmt.Println(string(s))
}

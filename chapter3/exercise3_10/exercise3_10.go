// Another version of "comma" using Buffer
// go run exercise3_10.go
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	// 1st part of the string
	first := n % 3
	if first == 0 {
		first = 3
	}
	buf.WriteString(s[:first])
	// other parts, each with 3 digits
	for i := first; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	tests := []string{
		"123",
		"1234567",
	}
	for _, t := range tests {
		fmt.Printf("%s -> %s\n", t, comma(t))
	}
}

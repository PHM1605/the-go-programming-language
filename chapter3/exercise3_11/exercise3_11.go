// comma() version that can handles +/- and '.' in floating point numbers
// go run exercise3_11.go
package main

import (
	"fmt"
	"strings"
)

// addComma is our basic "comma" function
func addComma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return addComma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

// powerful comma(): handle sign +/- and floating point
func comma(s string) string {
	// handle sign
	sign := ""
	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		s = s[1:]
	}
	// Split integer and fractional part
	intPart := s
	fracPart := ""
	if dot := strings.Index(s, "."); dot != -1 {
		intPart = s[:dot]
		fracPart = s[dot:] // keep the dot
	}
	// Add comma to integer part
	intPart = addComma(intPart)
	return sign + intPart + fracPart
}

func main() {
	tests := []string{
		"1234567",
		"-1234567",
		"+1234567",
		"1234567.89",
		"-1234567.8912",
		"123",
	}
	for _, t := range tests {
		fmt.Printf("%q -> %q\n", t, comma(t))
	}
}

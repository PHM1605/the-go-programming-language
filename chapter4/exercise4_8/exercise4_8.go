// Counts Unicode categories (letters, digits, spaces, punctuaions, symbols, others)
// go run exercise4_8.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// open text file
	file, err := os.Open("exercise4_8.txt")
	if err != nil {
		fmt.Println("open file", err)
		return
	}
	defer file.Close()

	// start counting categories
	categories := make(map[string]int)
	var invalid int                 // invalid rune count
	var utflen [utf8.UTFMax + 1]int // histogram of nbytes of runes

	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// invalid rune
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// ok rune, start counting and updating histogram
		utflen[n]++
		switch {
		case unicode.IsLetter(r):
			categories["letter"]++
		case unicode.IsDigit(r):
			categories["digit"]++
		case unicode.IsSpace(r):
			categories["space"]++
		case unicode.IsPunct(r):
			categories["punct"]++
		case unicode.IsSymbol(r):
			categories["symbol"]++
		default:
			categories["other"]++
		}
	}

	// Print category count
	fmt.Println("Category\tCount")
	for cat, n := range categories {
		fmt.Printf("%s\t%d\n", cat, n)
	}
	// print histogram
	fmt.Println("Nbytes of runes histogram")
	for i := 1; i < len(utflen); i++ {
		fmt.Printf("%d\t%d\n", i, utflen[i])
	}
	// print invalid runes
	if invalid > 0 {
		fmt.Printf("\nInvalid UTF-8 characters: %d\n", invalid)
	}
}

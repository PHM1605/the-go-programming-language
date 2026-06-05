// Compute counts of Unicode characters AND rune's length histogram
// go run charcount.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	invalid := 0
	// nbytes of runes histogram; utflen[0] never happens
	var utflen [utf8.UTFMax + 1]int

	// open file
	file, err := os.Open("charcount.txt")
	if err != nil {
		fmt.Println("open file: ", err)
		return
	}
	defer file.Close()

	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		} // run all file already
		// error in reading rune
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// invalid rune character => count "invalid" and onto next rune
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// okay rune
		counts[r]++
		utflen[n]++
	}

	// print result
	fmt.Printf("rune\tcount\n")
	// print rune counts
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	// print histogram of rune's nbytes
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	// print number of invalid runes
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

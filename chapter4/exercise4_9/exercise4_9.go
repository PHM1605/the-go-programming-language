// Count words in a text
// go run exercise4_9.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open text file
	file, err := os.Open("exercise4_9.txt")
	if err != nil {
		fmt.Println("open file: ", err)
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	input := bufio.NewScanner(file)
	// so that input.Text() returns a "word" instead of "line"
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	// print result
	for word, n := range counts {
		fmt.Printf("%s\t%d\n", word, n)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

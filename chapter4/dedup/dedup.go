// read a sequence of lines; print the first occurence of each distinct line
// go run dedup.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings to store distinct lines

	// to read from Stdin (but we will read from file for better convenience)
	// input := bufio.NewScanner(os.Stdin)

	file, err := os.Open("dedup.txt")
	if err != nil {
		fmt.Println("open file: ", err)
		return
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	// for NewScanner we must check error status
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

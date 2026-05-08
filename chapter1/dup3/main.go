// Dup3: count #occurences of lines
// read all lines in ONE file at once
// $go run main.go file1.txt file2.txt

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename) // read all lines in file
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// read each line
		// Note: file ends with "new line" will count blank line
		content := strings.TrimSpace(string(data)) // ... so we remove last line first
		for _, line := range strings.Split(content, "\n") {
			counts[line]++
		}
	}
	// print result
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

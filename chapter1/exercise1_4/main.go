// Dup2: counts occurences of lines; print file name as well
// go run main.go file1.txt file2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// NEW: map a "line-content" to "list of occuring files"
	lineFiles := make(map[string][]string)

	// read from stdin if no file from stdin's arguments
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles)
			f.Close()
		}
	}
	// print dict result
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, lineFiles[line]) // print list of occuring files as well
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		// {"hello": ["file1.txt", "file3.txt"]}
		lineFiles[line] = append(lineFiles[line], f.Name())
	}
}

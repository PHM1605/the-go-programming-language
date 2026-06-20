package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

// this is interface of "os.Write" (1st parameter of Fprint)
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	// running scan for words
	for scanner.Scan() {
		*c++
	}
	return len(p), scanner.Err()
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	// running scan for lines
	for scanner.Scan() {
		*c++
	}
	return len(p), scanner.Err()
}

func main() {
	text := `hello world
		this is golang
		nice day`

	var bc ByteCounter
	var wc WordCounter
	var lc LineCounter

	fmt.Fprint(&bc, text)
	fmt.Fprint(&wc, text)
	fmt.Fprint(&lc, text)

	fmt.Println("Bytes: ", bc)
	fmt.Println("Words: ", wc)
	fmt.Println("Lines: ", lc)
}

// Fetch several URLs concurrent and give summary table ONLY
// $ go build exercise1_10.go
// $ ./exercise1_10 https://www.reddit.com
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // create a "channel" of strings

	for i, url := range os.Args[1:] {
		filename := fmt.Sprintf("output%d.txt", i+1)
		go fetch(url, filename, ch)
	}

	// main() receives anything from "channel"
	for range os.Args[1:] {
		// print to terminal
		fmt.Println(<-ch) // receive from channel "ch"
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, filename string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send error to channel
		return
	}
	defer resp.Body.Close()

	// NEW: write content to files
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("creating %s: %v", filename, err)
		return
	}
	defer file.Close()

	// copy content to file
	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	// send result to channel
	ch <- fmt.Sprintf("%.2fs %7d bytes %s -> %s", secs, nbytes, url, filename)
}

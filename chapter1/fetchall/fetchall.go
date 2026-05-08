// Fetch several URLs concurrent and give summary table ONLY
// $ go build fetchall.go
// $ ./fetchall https://golang.org http://gopl.io https://godoc.org
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
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	// main() receives anything from "channel"
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel "ch"
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send error to channel
		return
	}
	// discard body but count bytes read
	nbytes, err := io.Copy(io.Discard, resp.Body)
	// clean up
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send error to channel
		return
	}
	secs := time.Since(start).Seconds()
	// send result to channel
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

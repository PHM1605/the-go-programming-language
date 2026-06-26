package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// NEW: request with cancellation capability
func fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	// send that request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func mirroredFetch(urls []string) string {
	// NEW: cancellation context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // will disable channel inside "ctx" (ctx.Done(), so <-ctx.Done() is called)

	// channel to receive and process our fetch
	responses := make(chan string, len(urls))
	for _, url := range urls {
		// each url has a goroutine
		go func(url string) {
			result, err := fetch(ctx, url)
			if err != nil {
				return
			}
			// NEW: to get response AND keep track of cancellation
			select {
			case responses <- result:
				// do nothing; but we get responses already, pumping to channel "responses"
			case <-ctx.Done():
				// do nothing
			}
		}(url)
	}

	// 1st response from the FASTEST url will come here and cancel the rest
	result := <-responses
	cancel()
	return result
}

func main() {
	urls := []string{
		"https://example.com",
		"https://go.dev",
		"https://golang.org",
	}
	start := time.Now()
	result := mirroredFetch(urls)
	fmt.Println(result)
	fmt.Println("elapsed: ", time.Since(start))
}

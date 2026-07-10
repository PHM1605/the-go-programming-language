package exercise93

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

// this heavy called will be cached
// NEW: "done" channel to signal cancellation of call by "close(done)"
func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// NEW: request that can be cancelled after sending "Request with context"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req = req.WithContext(ctx)

	// NEW: side goroutine to check status of "done" and "context"
	go func() {
		select {
		case <-done:
			cancel() // cancel context
		case <-ctx.Done(): // means the request has been finished normally
		}
	}()

	// sending request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body) // will return "byte[]" and "error"
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody) // create Memo object with function "func"
	defer m.Close()       // close "requests" channel inside Memo

	imcomingURLs := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://golang.org",
	}

	// use concurrent
	var n sync.WaitGroup
	// NEW: channel for early cancellation
	done := make(chan struct{})

	for _, url := range imcomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte))) // ([]byte): type assertion
			n.Done()
		}(url)
	}
	n.Wait()
}

func TestCancellation(t *testing.T) {
	m := New(httpGetBody) // create Memo object with function "func"
	defer m.Close()       // close "requests" channel in Memo

	const url = "https://golang.org"

	// call that cancels after 2 microsecond as SIDE goroutine
	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Millisecond)
		close(done)
	}()
	// real goroutine
	_, err := m.Get(url, done)
	if err == nil {
		t.Fatal("expected error from cancelled Get, got nil")
	}

	// call2: normal call
	done2 := make(chan struct{})
	value, err := m.Get(url, done2)
	if err != nil {
		t.Fatalf("expected uncancelled Get to succeed, got error: %v", err)
	}
	if len(value.([]byte)) == 0 {
		t.Fatal("expected non-empty body on uncancelled retry")
	}
	t.Logf("retry succeeded with %d bytes", len(value.([]byte)))
}

package memo1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

// this heavy called will be cached
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body) // will return "byte[]" and "error"
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody) // create Memo object with function "func" and "cache"
	imcomingURLs := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"golang.org",
	}

	// use concurrent
	var n sync.WaitGroup

	for _, url := range imcomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte))) // ([]byte): type assertion
			n.Done()
		}(url)
	}
	n.Wait()
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// return <name> and <length> of local file
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	// resp.Request: GET method; URL: https://example.com/images/cat.jpg
	// Path: images/cat.jpg OR "/" (if https://example.com/ only)
	// local: "cat.jpg" OR "/" => "index.html"
	local := path.Base(resp.Request.URL.Path) // /users/
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	// NOTE: we do NOT "defer f.Close()" here because it can cause error too (far from perfect)
	// NEW: Correct way to use "defer()"
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr // update "err" from Closing file before return (still, we prefer Copy's error below)
		}
	}()
	// copy data to file
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}

	filename, n, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("saved %s (%d bytes)\n", filename, n)
}

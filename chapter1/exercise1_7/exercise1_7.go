// Prints the content found at URL
// NEW: use io.Copy(dst, src) to NOT require a buffer large enough to hold entire body
// $ go build exercise1_7.go
// $ ./exercise1_7 http://gopl.io
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// read body of response
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1) // exit main() and show Error code 1
		}
	}
}

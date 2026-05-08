// Prints the content found at URL
// NEW: print response status as well
// $ go build exercise1_9.go
// $ ./exercise1_9 http://gopl.io
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
		// NEW: print response status first
		fmt.Println("Status: ", resp.Status)

		// read body of response
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1) // exit main() and show Error code 1
		}

		fmt.Printf("%s", b)
	}
}

package main

import (
	"chapter13/bzip"
	"io"
	"log"
	"os"
)

// command to run "./bzipper < /usr/share/dict/words | wc -c"
// means: Stdin is the /dict/words long file; Stdout is to "wc -c"
func main() {
	// this WriteCloser will replace Stdout (to compress data prior)
	w := bzip.NewWriter(os.Stdout)
	// io.Copy() will call "w.Write(<data-from-os.Stdin>)"
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

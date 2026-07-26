package main

import (
	"chapter13/exercise13_4/bzipgo"
	"io"
	"log"
	"os"
)

func main() {
	w := bzipgo.NewWriter(os.Stdout)
	defer w.Close()

	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

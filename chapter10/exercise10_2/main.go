package main

import (
	"chapter10/exercise10_2/archive"
	_ "chapter10/exercise10_2/archive/tar" // register tar in init()
	_ "chapter10/exercise10_2/archive/zip" // register zip in init()
	"flag"
	"fmt"
	"os"
)

var format = flag.String("format", "zip", "archive format")

func main() {
	flag.Parse()

	files, err := archive.ReadArchive(os.Stdin, *format)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, f := range files {
		fmt.Printf("%-30s %10d bytes\n", f.Name, f.Size)
	}

}

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"  // or blank import
	"image/jpeg" // or blank import
	"image/png"  // or blank import
	"io"
	"os"
)

var format = flag.String("format", "jpeg", "output format") // NEW: pointer to String

func convert(in io.Reader, out io.Writer) error {
	// NOTE: Decode() needs "blank import" to setup its config table
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	// which output format based on flag
	switch *format {
	case "jpeg", "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	default:
		return fmt.Errorf("unsupported output format: %s", *format)
	}
}

func main() {
	flag.Parse() // parse input flags from cmd

	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // NOTE: we need this to run "init()" so that "Decode()" can recognize PNG
	"io"
	"os"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in) // decode input image (*.png) to binary
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

// receive input from the "mandelbrot program"
// output tp stdout => to ">xxx.jpg" later
func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

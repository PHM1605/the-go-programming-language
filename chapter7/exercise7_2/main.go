package main

import (
	"fmt"
	"io"
	"os"
)

// new Writer class; like normal io.Writer but with counting bytes being pushed to it (hence "n")
type writerWrapper struct {
	w     io.Writer
	count int64
}

func (wW *writerWrapper) Write(p []byte) (int, error) {
	n, err := wW.w.Write(p)
	wW.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &writerWrapper{w: w}
	return cw, &cw.count
}

func main() {
	w, count := CountingWriter(os.Stdout)
	fmt.Fprint(w, "hello")
	fmt.Fprint(w, "world\n")
	fmt.Println("Bytes written = ", *count)
}

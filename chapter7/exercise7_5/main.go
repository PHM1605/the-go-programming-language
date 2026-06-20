package main

import (
	"fmt"
	"io"
	"strings"
)

// My new limit reader
type limitReader struct {
	r io.Reader
	n int64 // quota of reading
}

// to be compromise to "io.Reader" template
func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	// if what to read > remaining quota => read quota left only (e.g. 5>3=>3)
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}
	n, err := lr.r.Read(p) // read from normal Reader
	lr.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		r: r,
		n: n,
	}
}

func main() {
	r := strings.NewReader("Hello World")
	lr := LimitReader(r, 5) // our quota of reading
	buf := make([]byte, 100)

	n, err := lr.Read(buf) // read only 5 bytes => "Hello" only
	fmt.Printf("%q\n", buf[:n])

	// quota already = 0
	n, err = lr.Read(buf)
	fmt.Printf("%q\n", buf[:n])
	fmt.Println(err)
}

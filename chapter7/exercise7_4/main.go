package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// return something that satisfies "io.Reader" interface i.e. has "Read()" method
func NewReader(s string) *MyReader {
	return &MyReader{s: s}
}

type MyReader struct {
	s   string // immutable source e.g. `<html><head>.....</html>`
	pos int    // current position of our Reader
}

// to satisfy "io.Reader"
// <p>: to store reading result
// <r>: long text we must read and it's position
func (r *MyReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.pos:])
	r.pos += n
	return n, nil
}

func main() {
	htmlText := `
		<html>
			<body>
				<h1>Hello</h1>
			</body>
		</html>
	`
	doc, err := html.Parse(NewReader(htmlText))
	if err != nil {
		panic(err)
	}
	fmt.Println("Root node: ", doc.Type)
	fmt.Println(doc.FirstChild.Data)
}

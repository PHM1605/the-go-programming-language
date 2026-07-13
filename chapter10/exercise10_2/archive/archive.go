package archive

import (
	"fmt"
	"io"
)

// describe 1 file within the archive
type File struct {
	Name string
	Size int64
}

// receive a data stream from tar/zip file; output a list of Files
type ReaderFunc func(io.Reader) ([]File, error)

// map from a file type (e.g. "zip") to its corresponding "[]File releasor"
var formats = make(map[string]ReaderFunc)

// call this to update map above
func RegisterFormat(format string, fn ReaderFunc) {
	formats[format] = fn
}

// use map; get a list of Files from data-stream ("r") and its format-string (e.g. "zip")
func ReadArchive(r io.Reader, format string) ([]File, error) {
	fn, ok := formats[format]
	if !ok {
		return nil, fmt.Errorf("unknown archive format: %s", format)
	}
	return fn(r)
}

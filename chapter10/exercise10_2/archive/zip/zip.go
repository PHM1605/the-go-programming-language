package zip

import (
	"archive/zip"
	"bytes"
	"chapter10/exercise10_2/archive"
	"io"
)

// readZip: return list of files' metadata
// r: data stream
// File: {"name", "size"}
func readZip(r io.Reader) ([]archive.File, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	// get ZipReader from bytes of data-stream
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	// return files meta data from ZipReader
	var files []archive.File
	for _, f := range zr.File {
		files = append(files, archive.File{
			Name: f.Name,
			Size: int64(f.UncompressedSize64),
		})
	}

	return files, nil
}

func init() {
	archive.RegisterFormat("zip", readZip)
}

package tar

import (
	"archive/tar"
	"chapter10/exercise10_2/archive"
	"io"
)

// readTar: return list of files' metadata
// r: data stream
func readTar(r io.Reader) ([]archive.File, error) {
	// get TarReader
	tr := tar.NewReader(r)
	// loop over tar to get files' metadata
	var files []archive.File
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		files = append(files, archive.File{
			Name: hdr.Name,
			Size: hdr.Size,
		})
	}
	return files, nil
}

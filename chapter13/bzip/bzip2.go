package bzip

/*
#cgo LDFLAGS: -lbz2
#include <stdlib.h>
#include <bzlib.h>
int bz2compress(bz_stream *s, int action, char* in, unsigned* inlen, char* out, unsigned *outlen);
*/
import "C"
import (
	"io"
	"unsafe"
)

// to return "io.WriteCloser" interface, "writer" must implement "Write()" and "Close()" like below
type writer struct {
	w      io.Writer       // Stdout
	stream *C.bz_stream    // will wrap around "data []byte" of Write()
	outbuf [64 * 1024]byte // buffer before writing (decompressed) data to Stdout
}

// data: uncompressed data
// output: integer that shows "how many bytes from 'data' that were successfully consumed"
// compressed data: written to "writ.w" which is Stdout (as in 'bzipper')
func (writ *writer) Write(data []byte) (int, error) {
	if writ.stream == nil {
		panic("close")
	}
	var total int // #bytes from input that has been processed
	// stream Stdin until empty
	for len(data) > 0 {
		// inlen: input stream data length (meaning will be changed later)
		// outlen: output buffer length (meaning will be changed later)
		inlen, outlen := C.uint(len(data)), C.uint(cap(writ.outbuf))
		// C.BZ_RUN: action to perform (enum)
		C.bz2compress(
			writ.stream,
			C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&writ.outbuf)), &outlen,
		)
		// *inlen: changes meaning; now means "number-of-processed-bytes" (200)
		// *outlen: changes meaning; now means "number-of-bytes-being-produced" (100)
		total += int(inlen)
		data = data[inlen:]
		// outbuf[:outlen] now contains the compressed data
		if _, err := writ.w.Write(writ.outbuf[:outlen]); err != nil {
			return total, err
		}
	}

	return total, nil
}

func (writ *writer) Close() error {
	if writ.stream == nil {
		panic("closed")
	}
	// clean up C-stream
	defer func() {
		C.BZ2_bzCompressEnd(writ.stream) // release the stream buffers
		C.free(unsafe.Pointer(writ.stream))
		writ.stream = nil
	}()
	// draining "w.outbuf" py passing action=BZ_FINISH; empty input (nil, inlen=0)
	for {
		inlen, outlen := C.uint(0), C.uint(cap(writ.outbuf))
		r := C.bz2compress(writ.stream, C.BZ_FINISH, nil, &inlen, (*C.char)(unsafe.Pointer(&writ.outbuf)), &outlen)
		// put compressed data to Stdout
		if _, err := writ.w.Write(writ.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}

// out: Stdout i.e. input for "wc -c" in command pipeline
// => we wrap it with io.WriteCloser that compress data prior
// => must implement Write() and Close()
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	stream := (*C.bz_stream)(C.malloc(C.size_t(C.sizeof_bz_stream)))
	w := &writer{w: out, stream: stream}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

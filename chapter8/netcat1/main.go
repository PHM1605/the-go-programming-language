package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// create TCP connection
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// it hangs here to wait for server
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

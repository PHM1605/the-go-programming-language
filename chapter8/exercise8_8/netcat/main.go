package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	// hanging
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 2 goroutines here: 1 to send typing to Server; 1 to show from Server to screen
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

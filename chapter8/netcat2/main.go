package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	// it hangs here to wait for server
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

	// 2 goroutines here:
	// - 1 shows Server to Screen
	// - 1 sends from what Client has to Server
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

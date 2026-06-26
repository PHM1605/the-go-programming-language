package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// to "mark" from branch goroutine to main that we have received something from Server
	done := make(chan struct{})
	// branch goroutine (receiving echos)
	go func() {
		io.Copy(os.Stdout, conn) // hanging, print echos to screen
		log.Println("done")
		done <- struct{}{}
	}()

	// main goroutine: send what we type to server (hanging)
	mustCopy(conn, os.Stdin)
	// done typing => half close => send EOF to Server to Scanner()
	conn.(*net.TCPConn).CloseWrite()

	<-done // wait for branch goroutine to send something

	// completely close
	conn.Close()
}

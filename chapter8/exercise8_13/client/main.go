package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	// branch goroutine: hanging, receive what Server sends to us
	go func() {
		_, err := io.Copy(os.Stdout, conn) // hang, receive what Server sends to us
		if err != nil {
			log.Println("server disconnected:", err)
		}
		close(done) // send 0 to "<-done" statement
	}()

	// NEW: another goroutine to send what we type to Server
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Println("cannot send: ", err)
		}
	}()

	// wait for branch goroutine to finish too
	<-done
	log.Println("connection closed")
}

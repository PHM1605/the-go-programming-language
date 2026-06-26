package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	// it hangs here waiting for Server
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// NEW: create a channel from Client to Server
	done := make(chan struct{})
	// branch goroutine: hanging, receive what Server sends to us
	go func() {
		io.Copy(os.Stdout, conn) // hang, receive what Server sends to us
		log.Println("done")
		// NEW: write to the Channel an empty struct
		done <- struct{}{}
	}()

	// main goroutine: send what we type to Server
	mustCopy(conn, os.Stdin)
	conn.Close()
	// NEW: wait for branch goroutine to finish too
	<-done
}

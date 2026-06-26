package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second) // release time every 1s
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// when a Request reaches
		conn, err := listener.Accept() // it hangs here
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn) // BAD: handle ONE connection at a time
	}
}

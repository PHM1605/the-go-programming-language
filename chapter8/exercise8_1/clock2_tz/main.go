package main

import (
	"flag"
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
		// fire signal every 1s
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// to get "-port 8010"
	port := flag.String("port", "8000", "port number")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on localhost:%s", *port)

	for {
		// hangs here waiting
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // each connection also hangs here
	}
}

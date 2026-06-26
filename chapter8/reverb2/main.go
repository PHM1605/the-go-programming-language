package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// send back THREE times to Client with 3 different tones (1s apart)
		// NEW: gorountine even for only 1 Client (this Client)
		go echo(c, input.Text(), 1*time.Second)
	}
	if err := input.Err(); err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on localhost:8000")

	for {
		conn, err := listener.Accept() // hanging here
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

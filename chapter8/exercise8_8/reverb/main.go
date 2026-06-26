package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// handle each connection
func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1 of the 2 channels that we care for events (the other is time.After)
	lines := make(chan string)

	scanner := bufio.NewScanner(conn)
	// goroutine to scan what we'll receive from that connection
	go func() {
		for scanner.Scan() { // hanging
			lines <- scanner.Text()
		}
		// NOTE: it will prints this when TIMEOUT happens in time.After() and close connection (defer conn.Close())
		if err := scanner.Err(); err != nil {
			fmt.Println("read error: ", err)
		}
		close(lines)
	}()

	// main routine: handle BOTH what from client and what from timer
	for {
		select {
		// from user
		case line, ok := <-lines:
			if !ok {
				return // client disconnected
			}
			echo(conn, line, 1*time.Second)

		// from timer
		case <-time.After(10 * time.Second):
			fmt.Fprintln(conn, "Timeout: disconnected.")
			return
		}
	}
}

// print input word with decreasing tone; 1s apart
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Echo server listening on localhost:8000")

	for {
		// hang here waiting
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// handle each connection
		go handleConn(conn)
	}
}

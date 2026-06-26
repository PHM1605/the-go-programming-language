package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// echoing in decreasing tone; 1s apart
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	// hang here waiting for what Client will send
	// as each "echo" will take > 3s; each sentence from that ONE Client will be handled concurrent here
	var wg sync.WaitGroup // control #echo-processes
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		text := scanner.Text()
		// NEW: count up #goroutines here
		wg.Add(1)
		go func(msg string) {
			defer wg.Done() // count down #goroutines here
			echo(c, msg, 1*time.Second)
		}(text)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// NEW: here we wait for all echo processes to finish
	wg.Wait()
	// Close write-half only (will signal EOF for Client)
	if tcpConn, ok := c.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
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
		// hang here wating for requests
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// scan server responses concurrently
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			txt := scanner.Text()
			fmt.Println(txt)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}()

	// send commands to server
	scanner := bufio.NewScanner(os.Stdin)
	// hang here for typing commands => send to server
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

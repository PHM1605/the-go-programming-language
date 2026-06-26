package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		// NewYork=localhost:8010
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			log.Fatalf("bad argument: %s", arg)
		}
		name := parts[0]    // NewYork, London
		address := parts[1] // localhost8010, localhost:8020

		// send requests to those address concurrently
		go func(name, address string) {
			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			// scanning from a remote URL
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				fmt.Printf("%-10s %s\n", name, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Printf("%s: %v", name, err)
			}
		}(name, address)
	}
	// keep hanging main Go process
	select {}
}

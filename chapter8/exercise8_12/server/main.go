package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// for each client: store its "channel" and "name" of that client
type client struct {
	ch   chan string
	name string
}

var (
	entering = make(chan client) // NEW: channel to send/receive outgoing clients
	leaving  = make(chan client) // NEW: channel to send/receive outgoing clients
	messages = make(chan string) // all incoming clients' messages
)

func broadcaster() {
	// set of clients
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// broadcasting incoming messages to all receiving-strings channels
			for cli := range clients {
				cli.ch <- msg
			}
		// which Client is entering => add to the set "clients"
		case cli := <-entering:
			clients[cli] = true
			// names of current Clients
			var names []string
			for c := range clients {
				names = append(names, c.name)
			}
			// NEW: print current Clients
			cli.ch <- "Current users:"
			for _, name := range names {
				cli.ch <- " - " + name
			}
		// which channel is leaving => remove from the set "clients" & close that channel
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// hanging; to push strings that enter (cli := <-msg) from other users to this connection
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleConn(conn net.Conn) {
	// identity of connection
	who := conn.RemoteAddr().String()
	// channel to push messages to this client (owner of "conn")
	ch := make(chan string)
	// NEW: now we store the whole Client instead of channel of that Client only
	cli := client{
		ch:   ch,
		name: who,
	}
	go clientWriter(conn, ch) // hanging inside to keep channel alive to send strings to Client

	ch <- "You are " + who           // send to Client "You are xyz"
	messages <- who + " has arrived" // send to messages pool "xyz has arrived" (to be sent to other Clients)
	entering <- cli                  // to update broadcaster's clients pool

	// scanning what this client is sending
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // hanging
		messages <- who + ": " + scanner.Text() // send to messages pool to be sent to all other Clients
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// that Client sends EOF
	leaving <- cli                // remove from broadcaster's clients pool //
	messages <- who + " has left" // send to messages pool to be sent to all other Clients

	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// NEW: like a channels manager
	go broadcaster()

	for {
		conn, err := listener.Accept() // hanging
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // per-client goroutine
	}
}

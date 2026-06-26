package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("LIFT OFF")
}

func main() {
	// create abort channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // hanging; read 1 byte from keyboard
		abort <- struct{}{}            // send that byte to channel "abort", wait for receiving
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	// time.After(t): returns a channel; inside starts a goroutine (go func(){...}() with delay "t" seconds inside)
	// NOTE: when 1 event comes to 1 of the TWO OR MORE CHANNELS
	select {
	case <-time.After(10 * time.Second): // somethings comes to this channel automatically after 10s
		// Do nothing
	case <-abort: // something comes to channel "abort"
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

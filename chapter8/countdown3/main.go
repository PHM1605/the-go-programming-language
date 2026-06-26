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
	// time.Tick(t): returns a channel
	// inside starts a goroutine (go func(){...}()) that injects into that channel every "t" seconds
	tick := time.Tick(1 * time.Second) // NOTE: goroutine LEAK here
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// control 2 channels here
		select {
		case <-tick:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

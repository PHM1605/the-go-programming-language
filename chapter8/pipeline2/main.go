package main

import "fmt"

func main() {
	// create 2 channels
	naturals := make(chan int)
	squares := make(chan int)

	// Counter goroutine: send 0,1,2,3... to naturals channel
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer goroutine: receive 0,1,2.. from naturals, send 0,2,4,.. to squares channel
	go func() {
		// NEW syntax: x <- naturals but stop if flag "ok" is false (implicitly in "range" operator)
		for x := range naturals {
			squares <- x * x // send to "squares" channel
		}
		close(squares)
	}()
	// Main Printer goroutine: Receive from Squarer and print
	for x := range squares {
		fmt.Println(x)
	}
}

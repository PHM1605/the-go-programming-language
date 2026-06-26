package main

import "fmt"

func main() {
	// create 2 channels
	naturals := make(chan int)
	squares := make(chan int)

	// Counter goroutine: send 0,1,2,3... to naturals channel
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()
	// Squarer goroutine: receive 0,1,2.. from naturals, send 0,2,4,.. to squares channel
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// Main goroutine: Receive from Squarer and print
	for {
		fmt.Println(<-squares)
	}
}

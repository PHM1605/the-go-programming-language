package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	// this will run forever UNTIL another go process reaches end of program
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // \r means "carriage return"
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow process
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

package main

import "fmt"

// send 0,1,2..
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out) // NOTE: close() only for "out-only-channels"
}

// receive 0,1,2.. and send 0,1,4..
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// receive 0,1,4.. and print
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 3 goroutines
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

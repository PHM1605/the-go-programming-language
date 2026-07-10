package main

import "fmt"

func main() {
	// 2 goroutines, 1 prints 0s, 1 prints 1s
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

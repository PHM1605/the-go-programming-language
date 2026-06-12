package main

import "fmt"

// print integer in descending order until "panic"
// then it releases all DEFERs in reverse order
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(3)
}

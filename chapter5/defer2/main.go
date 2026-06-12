package main

import (
	"fmt"
	"os"
	"runtime"
)

// print integer in descending order until "panic"
// then it releases all DEFERs in reverse order
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	// buf: where to write the stack trace
	// false: whether to include all goroutines (in this case FALSE => only current goroutine)
	// n: how many bytes of the current stack trace is relevant
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n]) // print to screen the stack trace
}

// we will see f(3)->f(2)->f(1)
// panic so -> defer 1 -> defer 2 -> defer 3 -> printStack()
func main() {
	defer printStack()
	f(3)
}

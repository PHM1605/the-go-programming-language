package main

import "fmt"

// squares() is a function-call => yield anonymous function "func() int"
// x is property OF squares()
func squares() func() int {
	var x int // NOTE: cached value here
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()   // will allocate "x" AND return an x-attached-anonymous-function "func() int"
	fmt.Println(f()) // f implicitly has "x" inside
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

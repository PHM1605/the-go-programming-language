// Find n-th Fibonacci number
package main

import "fmt"

func main() {
	n := 6
	fmt.Printf("%dth-Fibonacci-number is %d\n", n, fib(n))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

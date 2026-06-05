// Demonstrate how "append()" works in slice
// go run append.go
package main

import "fmt"

// ... means "variadic" = "any number of ints"
func appendInt(x []int, y ...int) []int {
	var z []int // what we should return
	zlen := len(x) + len(y)
	// if capacity of x is enough
	if zlen <= cap(x) {
		z = x[:zlen] // extend <z> to have enough space for new element being added too
	} else {
		// if capacity of x is not yet enough; x has capacity of 2 elements, z needs 3
		zcap := zlen
		// z capacity then DOUBLED to 6
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		// allocate slice=3 (1 more for new element) and capacity of 6
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y) // copy all elements of y into z (expanded)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

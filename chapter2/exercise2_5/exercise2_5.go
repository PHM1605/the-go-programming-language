// Count number of 1 bits from an "uint64" ("population count")
// NEW: using the fact that x&(x-1) turn off the last "1" bit to "0"
package main

import "fmt"

// how many 1s for 8-byte numbers
func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func main() {
	fmt.Printf("%d\n", PopCount(255)) // should be 8 bits of "1"
}

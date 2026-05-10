// Count number of 1 bits from an "uint64" ("population count")
// NEW: using "for" loop and WITHOUT lookup table
package main

import "fmt"

// how many 1s for 8-byte numbers
func PopCount(x uint64) int {
	var res int
	for i := 0; i < 64; i++ {
		res += int(byte(x>>i) & 1)
	}
	return res
}

func main() {
	fmt.Printf("%d\n", PopCount(1025)) // should be 2 bits of "1"
}

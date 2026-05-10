// Count number of 1 bits from an "uint64" ("population count")
package main

import "fmt"

// pc[i] is the population count of "i"
// e.g. pc[8] is number of "1" bits of 8 i.e. 1 ON bit
var pc [256]byte

// numbers from 0->255 (1 byte) has how many 1s each
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// how many 1s for 8-byte numbers
// use 1-byte-lookup-table 8 times
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func main() {
	fmt.Printf("%d\n", PopCount(1025)) // should be 2 bits of "1"
}

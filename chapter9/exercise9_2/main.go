package main

import (
	"fmt"
	"sync"
)

// array of 256 byte-elements; population count of 0, 1, 2, 3...
// pc[8] is number of "1"s bits of the number value "8" i.e. 1 ON bit
var pc [256]byte

// NEW: boolean-mutex
var once sync.Once

func initPopulationCount() {
	// i = index = 0..255
	for i := range pc {
		//if i even => byte(i&1)=0; if i odd => byte(i&1)=1
		// i/2 means "move all pattern to right 1"
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// how many 1s in a 8-byte numbers
// shifting 1-byte-lookup-table 8 times
func PopCount(x uint64) int {
	// NEW: we initialize ONCE with MUTEX
	once.Do(initPopulationCount)

	return int(
		pc[byte(x>>(0*8))] + // byte() takes last byte; then p[<byte value>] gets the #1s
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func main() {
	// NEW: to wait for all goroutines to finish before done
	var wg sync.WaitGroup

	numbers := []uint64{1025, 7, 255, 1024, 999999}

	for _, n := range numbers {
		wg.Add(1)

		go func(x uint64) {
			defer wg.Done()
			fmt.Printf("PopCount(%d) = %d\n", x, PopCount(x))
		}(n)
	}
	wg.Wait() // wait for all goroutines here
}

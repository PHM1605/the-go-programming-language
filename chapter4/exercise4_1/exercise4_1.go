// Write a function that counts the number of bits that are different in two SHA256 hashes (see PopCount() from Section 2.6.2)
// go run exercise4_1.go
package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

// each byte is 2 chars e.g. "2c"
func CountDiffBits(a, b [32]byte) int {
	count := 0
	for i := range a {
		count += bits.OnesCount8(a[i] ^ b[i])
	}
	return count
}

func main() {
	h1 := sha256.Sum256([]byte("hello"))
	h2 := sha256.Sum256([]byte("world"))
	// each hash is 256 bits = 32 bytes with each byte is 2 chars e.g. "2c"
	fmt.Printf("h1: %x\n", h1)
	fmt.Printf("h2: %x\n", h2)
	fmt.Printf("Different bits: %d / 256\n", CountDiffBits(h1, h2))
}

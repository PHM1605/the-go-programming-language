package main

import (
	"fmt"
	"sync"
)

// binary semaphore: ensure count of 1 goroutine access "balance" at the same time
var (
	mu      sync.Mutex
	balance int // our main variable
)

// NEW: programmer make sure INTERNAL FUNCTION "deposit" (lower-case) always wrapped in "lock" & "unlock" before use
// this INTERNAL function "deposit" NEVER has "lock"/"unlock"
func deposit(amount int) { balance += amount }
func Deposit(amount int) {
	// until "Unlock" is freed "balance" is safe inside
	mu.Lock()
	defer mu.Unlock()

	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()

	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func main() {
	Deposit(100)
	Deposit(20)

	Withdraw(10)

	fmt.Println(Balance()) // 110
}

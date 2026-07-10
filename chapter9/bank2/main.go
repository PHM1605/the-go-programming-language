package main

import "fmt"

// binary semaphore: ensure count of 1 goroutine access "balance" at the same time
var (
	sema    = make(chan struct{}, 1) // binary semaphore
	balance int                      // our main variable
)

func Deposit(amount int) {
	// until "sema" is freed "balance" is safe inside
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}
func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {
	Deposit(100)
	Deposit(20)

	fmt.Println(Balance()) // 120
}

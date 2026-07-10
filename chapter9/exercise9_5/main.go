package main

import (
	"fmt"
	"time"
)

func main() {
	// number of back-and-forth cycles
	const N = 3000000
	ping := make(chan struct{})
	pong := make(chan struct{})
	// done: to mark when N echo-processes finish AND record time
	done := make(chan time.Duration)

	// goroutine A: initiate each round-trip
	go func() {
		start := time.Now()
		for i := 0; i <= N; i++ {
			ping <- struct{}{}
			<-pong // hanging to wait for "pong's response"
		}
		// when everything finishes
		done <- time.Since(start)
	}()

	// goroutine B: echo back
	go func() {
		for i := 0; i <= N; i++ {
			// get what A sends
			<-ping
			// echo back
			pong <- struct{}{}
		}
	}()

	elapsed := <-done

	totalComms := 2 * N
	fmt.Printf("%d round trips = %d channel communications\n", N, totalComms)
	fmt.Printf("elapsed: %v\n", elapsed)
	fmt.Printf("%.0f communications/sec\n", float64(totalComms)/elapsed.Seconds())
}

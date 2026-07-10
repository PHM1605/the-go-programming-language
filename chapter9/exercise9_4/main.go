package main

import (
	"fmt"
	"runtime"
	"time"
)

// hanging; forwarding an integer from "in" channel to "out" channel
func stage(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v
	}
	// exit when "in" was closed
	close(out)
}

func main() {
	const N = 100000

	// read memory at the beginning
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	// build a lot of goroutines; one after another
	// goroutine 0: first -> next(i=0)
	// goroutine 1: next(i=0) -> next(i=1)
	// ...
	first := make(chan int)
	cur := first
	buildStart := time.Now()
	for i := 0; i < N; i++ {
		next := make(chan int)
		go stage(cur, next)
		cur = next
	}
	last := cur
	buildTime := time.Since(buildStart)

	// read memory at the end of the building process
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// send a value through all goroutines
	transitStart := time.Now()
	first <- 42
	v := <-last
	transitTime := time.Since(transitStart)

	// close the first => will close next(i=0) in goroutine 0 => will close next(i=1) in goroutine 1 etc.
	close(first)

	fmt.Printf("stages=%d value_out=%d\n", N, v)
	fmt.Printf("build time=%v\n", buildTime)
	fmt.Printf("transit_time=%v\n", transitTime)
	fmt.Printf("heap takes=%.0f MB\n", float64(m1.HeapAlloc-m0.HeapAlloc)/1e6)
}

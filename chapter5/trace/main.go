package main

import (
	"log"
	"time"
)

// trace() body runs IMMEDIATELY at first
// the "return" function will be called at the end by "defer"
func trace(msg string) func() {
	// this is onStart() log time
	start := time.Now()
	log.Printf("Start time %s of message %q\n", start, msg)
	// this is onExit() log time
	return func() {
		log.Printf("exit %s after %s\n", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	// NOTE: we log BOTH time start and time end with ONE call
	defer trace("bigSlowOperation")() // notice the parentheses at the end
	// Simulate a long operation
	time.Sleep(10 * time.Second)
}

func main() {
	bigSlowOperation()
}

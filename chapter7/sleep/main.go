package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.Duration() is a function that returns a "*time.Duration" type object
// flag is "-period", default of "1 second"
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period) // *period => time.Duration type, with "String()" method
	time.Sleep(*period)
	fmt.Println()
}

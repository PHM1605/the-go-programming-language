package main

import (
	"flag"
	"fmt"
	"tempflag/tempconv"
)

// function that returns a <*Celsius> variable
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

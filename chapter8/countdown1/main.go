package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("LIFT OFF")
}

func main() {
	fmt.Println("Commencing countdown.")
	// this channel will receive Events every 1s
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick // it will pause this receive untils an Event is sent
	}
	launch()
}

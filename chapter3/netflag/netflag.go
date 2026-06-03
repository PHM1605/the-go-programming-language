// Program to provide bitwise operations
// go run netflag.go
package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link'
	FlagMulticast                      // supports multicast access capability
)

// check if bit 0 is turned ON
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

// Turn off bit 0
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

// Turn ON bit 1, keep every other bits the same
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

// check if bit 1 OR bit 5 is ON (broadcast)
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp // 10001
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10001 true"
	TurnDown(&v)                         // turn off bit 0 of v
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10000 false"
	SetBroadcast(&v)                     // 10010
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v))  // "10010 true"
}

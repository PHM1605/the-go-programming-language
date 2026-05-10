// Find greatest common divisor (GCD)
package main

import "fmt"

func main() {
	a, b := 24, 9
	fmt.Printf("GCD of %d and %d is %v\n", a, b, gcd(a, b))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

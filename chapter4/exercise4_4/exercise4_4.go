// Write a version of "rotate" that operates in a single pass
// go run exercise4_4.go
package main

import "fmt"

func rotate(s []int, d int) {
	tmp := append(s[d:], s[:d]...)
	copy(s, tmp)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s) // {2,3,4,5,0,1}
}

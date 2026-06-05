// Write a reverse function that takes ARRAY POINTER instead of slice
// go run exercise4_3.go

package main

import "fmt"

func reverse(a *[6]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [...]int{4, 5, 6, 7, 8, 9}
	reverse(&a)
	fmt.Println(a)
}

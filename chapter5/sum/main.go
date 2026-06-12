package main

import "fmt"

// Sum of integers
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	// pass a slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

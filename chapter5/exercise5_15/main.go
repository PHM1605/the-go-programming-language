package main

import (
	"fmt"
)

// Make sure "max" has at least 1 argument
func max(vars ...int) int {
	if len(vars) == 0 {
		return 0
	}
	m := vars[0]
	for _, v := range vars[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// Make sure "min" has at least 1 argument
func min(vars ...int) int {
	if len(vars) == 0 {
		return 0
	}
	m := vars[0]
	for _, v := range vars[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func main() {
	vars := []int{3, 5, 1, 2}
	fmt.Println("Max is: ", max(vars...))
	fmt.Println("Min is: ", min(vars...))

	fmt.Println("Max of nothing: ", max())
	fmt.Println("Min of nothing: ", min())
}

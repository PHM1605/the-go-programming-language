// Tree sort structure
// go run treesort.go
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// add <value> into tree <t> in ascending order
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// return elements in Tree (as form []int)
// => [1,3,5,8]
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left) // if Left is null then return same <values>
		// NOTE: this is the main line => add current Node to <values>
		values = append(values, t.value)
		values = appendValues(values, t.right) // if Right is null then return same <values>
	}
	return values
}

// e.g. values: {5,3,8,1}
func Sort(values []int) {
	var root *tree
	// add a list of ints into tree, in ascending order
	for _, v := range values {
		root = add(root, v)
	}
	// values[:0]: create a slice of length=0 and capacity=4 (still point to original array)
	appendValues(values[:0], root)
}

func main() {
	values := []int{5, 3, 8, 1}
	fmt.Println("Before: ", values)
	Sort(values)
	fmt.Println("After: ", values)
}

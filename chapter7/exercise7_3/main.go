package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// add a number to Tree, return new Tree
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

// return list of integers in Tree after appending
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// build a Tree from a slice of ints
func buildTree(values []int) *tree {
	var root *tree
	// add a list of ints into tree, in ascending order
	for _, v := range values {
		root = add(root, v)
	}
	return root
}

// values: {5,3,8,1}
func Sort(values []int) {
	root := buildTree(values)
	// values[:0]: create a slice of length=0 and capacity of the old "values" i.e. =4
	appendValues(values[:0], root)
}

// NEW: for printing to screen
func (t *tree) String() string {
	// get slice of sorted ints from Tree <t>
	values := appendValues(nil, t)
	return fmt.Sprintf("Tree values: %v", values)
}

func main() {
	values := []int{5, 3, 8, 1}
	root := buildTree(values)
	fmt.Println(root) // implicitly use "String()" here
}

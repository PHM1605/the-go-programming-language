package main

import (
	"fmt"
)

// prereqs: map a course (string) to a list of prerequisite courses ([]string)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string            // what we return
	seen := make(map[string]bool) // flag which courses have we picked

	// visitAll() uses visitAll() inside => must declare as a variable first
	// items: now only 1 course OR it's list of prerequisite courses
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item]) // append list of prerequisites first
				order = append(order, item)
			}
		}
	}

	// NEW: iterate through map directly
	for course := range m {
		if !seen[course] {
			visitAll([]string{course}) // convert "xyz" => ["xyz"]
		}
	}

	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

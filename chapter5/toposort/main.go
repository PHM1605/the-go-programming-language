package main

import (
	"fmt"
	"sort"
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
	// items: sorted keys in our original <map> OR list of prerequisites
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

	// pick out keys of original <map> and sort in ascending order
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// actual sort; <order> will be appended here inside
	visitAll(keys)

	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

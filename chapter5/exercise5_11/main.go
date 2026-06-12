package main

import (
	"fmt"
	"log"
)

// prereqs: map a course (string) to a list of prerequisite courses ([]string)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},

	// NEW: internal cycle here
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string // what we return

	// NEW: we have THREE states now instead of "seen" boolean only
	const (
		unvisited = iota
		visiting
		done
	)
	state := make(map[string]int) // flag which courses have we have (not) visited OR visit twice

	// visit() uses visit() inside => must declare as a variable first
	// item: each key in original <map> OR one of the prerequisites
	var visit func(string) error
	visit = func(item string) error {
		// check state of each item
		switch state[item] {
		case done:
			return nil
		case visiting:
			return fmt.Errorf("cycle detected involving %q", item)
		}
		// NOTE: if we are going deeper of "calculus"
		// => go to "algebra" => see "calculus" again in "visiting" stage we know it's CYCLE
		state[item] = visiting
		for _, prereq := range m[item] {
			if err := visit(prereq); err != nil {
				return err
			}
		}
		// okay for that item => mark DONE and add to list
		state[item] = done
		order = append(order, item)
		return nil
	}

	// iterate through map directly
	for course := range m {
		if state[course] == unvisited {
			if err := visit(course); err != nil {
				return nil, err
			}
		}
	}
	return order, nil
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	// Display
	for i, course := range order {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

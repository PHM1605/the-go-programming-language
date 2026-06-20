package main

import (
	"fmt"
	"sort"
)

// sort.Interface: a "slice-type" must implement this to be used in "sort(abc)" function
// type Interface interface {
//   Len() int
//   Less(i, j int) bool
//   Swap(i, j int)
// }

// 1st "slice-type" to be used in sort(...)
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] } // to be in ascending order
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	names := []string{"abc", "xyz", "def"}
	sort.Sort(StringSlice(names)) // StringSlice(names) will create a reference of different type to underlying array - no copy it; hence we use the old one "names"
	fmt.Println(names)
	// lib built in
	// sort.Strings(names)
}

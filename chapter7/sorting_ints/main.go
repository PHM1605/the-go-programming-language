package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // false
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values)) // true

	// IntSlice: type from libraty; add methods being implemented for Interface to "[]int"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // [4,3,1,1]
	fmt.Println(sort.IntsAreSorted(values)) // false
}

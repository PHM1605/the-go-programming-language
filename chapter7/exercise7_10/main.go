package main

import (
	"fmt"
	"sort"
)

// must be complied to sort.Interface
type MySlice []int

func (x MySlice) Len() int           { return len(x) }
func (x MySlice) Less(i, j int) bool { return x[i] < x[j] }
func (x MySlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// s is Interface => has only Len(), Swap() and Less() functions
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - 1 - i
		// if front-element and back-element are equal
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false // not palindrome as soon as we see non-matched elements
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(MySlice{1, 2, 3, 2, 1})) // true
	fmt.Println(IsPalindrome(MySlice{1, 2, 3, 4, 5})) // false
}

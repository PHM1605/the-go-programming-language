// Write in-place function to eliminate adjacent duplicates in a []string slice
// go run exercise4_5.go
package main

import "fmt"

func uniq(s []string) []string {
	i := 1 // index for result
	for j := 1; j < len(s); j++ {
		if s[j] != s[j-1] {
			s[i] = s[j]
			i++
		}
	}
	return s[:i]
}

func main() {
	s := []string{"a", "a", "b", "b", "b", "c", "a", "a"}
	fmt.Println(uniq(s)) // {"a","b","c","a"}
}

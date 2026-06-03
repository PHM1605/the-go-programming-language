// test if one string is an ANAGRAM of another (e.g. "egg" is anagram of "geg")
// go run exercise3_12.go
package main

import "fmt"

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)
	// count characters in "a"
	for _, ch := range a {
		count[ch]++
	}
	// loop over "b"; subtract count of characters in "a"
	// => if any char count < 0 then not anagram
	for _, ch := range b {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		a, b string
	}{
		{"listen", "silent"},
		{"triangle", "integral"},
		{"apple", "papel"},
		{"rat", "car"},
		{"aabb", "bbaa"},
	}

	for _, t := range tests {
		fmt.Printf("%q & %q => %v\n", t.a, t.b, isAnagram(t.a, t.b))
	}
}

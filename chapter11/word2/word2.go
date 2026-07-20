package word

import "unicode"

// IsPalindrome reports where a string is the same forward and backward
func IsPalindrome(s string) bool {
	var letters []rune // slower
	// letters := make([]rune, 0, len(s)) // faster
	for _, r := range s {
		// take the runes first (exclude special chars)
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	// check palindrome from runes
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

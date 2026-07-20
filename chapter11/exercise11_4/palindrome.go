package exercise113

// IsPalindrome reports where a string is the same forward and backward
func IsPalindrome(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

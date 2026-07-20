package exercise113

import (
	"math/rand"
	"testing"
	"time"
)

// return a palindrome-string WITH spaces and punctuations
func randomPalindrome(rng *rand.Rand) string {
	// length of palin-string; choose random up to 24
	n := rng.Intn(25)
	runes := make([]rune, n) // string of n chars

	for i := 0; i < (n+1)/2; i++ {
		var r rune
		// 20% of using punctuations or spaces
		punct := []rune{' ', ',', '.', '!', '?', ':', ';', '\''}
		if rng.Intn(5) == 0 {
			r = punct[rng.Intn(len(punct))]
		} else {
			r = rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		}

		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// return a non-palindrome-string
// by create palindrome first and change 1 letter
func randomNonPalindrome(rng *rand.Rand) string {
	r := []rune(randomPalindrome(rng))
	// need at least 2 chars
	if len(r) < 2 {
		return "ab"
	}
	// change 1 letter at the begining
	r[0]++
	return string(r)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed)) // New returns pointer-to-Rand object
	// generate 1000 palin-strings to test
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	// random generate - create *Rand
	rng := rand.New(rand.NewSource(seed))
	// generate 1000 non-palindromes
	for i := 0; i < 1000; i++ {
		s := randomNonPalindrome(rng)
		if IsPalindrome(s) {
			t.Errorf("IsPalindrome(%q) = true", s)
		}
	}
}

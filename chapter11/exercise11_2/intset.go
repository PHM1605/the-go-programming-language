package exercise112

// rows of 64 bits each
// 1st row: for numbers 0->63
// 2nd row: for numbers 64->127 etc.
type IntSet struct {
	words []uint64
}

// check if IntSet has number "x"
func (s *IntSet) Has(x int) bool {
	// find which row and which column first
	word, bit := x/64, uint(x%64)
	// check if that bit at <row, col> == <word, bit> is ON or not
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// add a number to the Set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// add more rows (all 0s) if x too big
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit // turn that bit ON
}

// convert Set s to Set "s UNION t"
func (s *IntSet) UnionWith(t *IntSet) {
	// scan each row of t
	for i, tword := range t.words {
		// still in <s>'s scope
		if i < len(s.words) {
			s.words[i] |= tword
		} else { // t's scope => add that to s's scope
			s.words = append(s.words, tword)
		}
	}
}

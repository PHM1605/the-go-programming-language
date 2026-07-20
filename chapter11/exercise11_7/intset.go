package exercise117

// rows of 64 bits each
// representing existence of numbers 0->63 on 1st row; numbers 64->127 on 2nd row etc.
type IntSet struct {
	words []uint64
}

// check if Set has positive number "x"
func (s *IntSet) Has(x int) bool {
	// find that number x is which row and which column
	word, bit := x/64, uint(x%64)
	// check if bit at <row,col> is ON
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// add a number to the Set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// if we insert number 67 => word=1, bit=3
	// => we insert to Set 2 rows with values 0 all
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit // turn ON that bit
}

// convert set "s" to "s UNION t"
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

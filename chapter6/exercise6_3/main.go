package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

// rows of 64 bit each (each = 1 word)
// represent numbers from 0->63 for 1st row, then 64->127 for 2nd row, then etc.
// e.g. set has 2 elements {2, 68} then bit-2 on 1st row and bit-5 on 2nd row are ON
type IntSet struct {
	words []uint64
}

// function to check if Set has positive number "x"
func (s *IntSet) Has(x int) bool {
	// find which row and which column first
	word, bit := x/64, uint(x%64)
	// check if that bit at <row,col> == <word,bit> is turned ON
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// add a number to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// if we insert number e.g. 67 => word = 1, bit = 3
	// => we insert to SET 2 rows wit values 0 all
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit // turn ON that bit 3 of word 1 (i.e. bit-67)
}

// add many integers at once
func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

// return number of elements in a Set
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += bits.OnesCount64(word)
	}
	return count
}

// remove an element from Set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit // &^ means "AND NOT"
}

// copy a Set
func (s *IntSet) Copy() *IntSet {
	copyWords := make([]uint64, len(s.words))
	copy(copyWords, s.words)
	return &IntSet{
		words: copyWords,
	}
}

// remove all elements from a Set
func (s *IntSet) Clear() {
	s.words = nil
}

// convert set <s> to the "<s> OR <t>" (Union)
func (s *IntSet) UnionWith(t *IntSet) {
	// scan each row of <t>
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// <s> AND <t>
func (s *IntSet) IntersectWith(t *IntSet) {
	// scan over <s>
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// <s> DIFFERENCE <t> (elements in <s> but not in <t>)
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i] // &^ means "AND NOT"
		}
	}
}

// <s> XOR <t> (elements in <s> OR in <t> but NOT BOTH)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	// scan <t>
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword // ^ means "XOR"
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// print the Set
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		// if that row all ZERO bits => we skip, don't print
		if word == 0 {
			continue
		}
		// print each bit on each row
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				// print " " before each number (if not the 1st number)
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var a, b IntSet
	a.AddAll(1, 2, 3, 4)
	b.AddAll(3, 4, 5, 6)
	fmt.Println("a = ", a.String())
	fmt.Println("b = ", b.String())

	// Test intersection
	c := a.Copy()
	c.IntersectWith(&b)
	fmt.Println("Intersection result: ", c.String())

	// Test a DIFFERENCE WITH b (in <a> but not in <b>)
	d := a.Copy()
	d.DifferenceWith(&b)
	fmt.Println("a DIFFERENCE WITH b result: ", d.String())

	// Test XOR
	e := a.Copy()
	e.SymmetricDifference(&b)
	fmt.Println("a XOR b result: ", e.String())
}

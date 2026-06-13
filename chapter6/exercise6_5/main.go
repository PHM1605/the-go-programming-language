package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

// in 32-bit machine: (^uint(0) >> 63) means 0; then 32 << 0 == 32
// in 64-bit machine: (^uint(0) >> 63) means 1; then 32 << 1 == 64
const wordSize = 32 << (^uint(0) >> 63)

// rows of 32 or 64 bit each (each = 1 word)
// if set has 4,7 as elements => bit 4 and 7 on 1st row is ON
type IntSet struct {
	words []uint
}

// function to check if Set has positive number "x"
func (s *IntSet) Has(x int) bool {
	// find which row and which column first
	word, bit := x/wordSize, uint(x%wordSize)
	// check if that bit at <row,col> == <word,bit> is turned ON
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// add a number to the set
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
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
		count += bits.OnesCount(word)
	}
	return count
}

// remove an element from Set
func (s *IntSet) Remove(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit // &^ means "AND NOT"
}

// copy a Set
func (s *IntSet) Copy() *IntSet {
	copyWords := make([]uint, len(s.words))
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
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				// print " " before each number (if not the 1st number)
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// NEW: return slice of elements in that set
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, wordSize*i+j)
			}
		}
	}
	return elems
}

func main() {
	var s IntSet
	s.AddAll(1, 9, 42, 123)
	fmt.Println("Elements of s: ")
	for _, v := range s.Elems() {
		fmt.Println(v)
	}
}

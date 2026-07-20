package main

import (
	"bytes"
	"fmt"
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
	// => we insert to SET 2 rows with values 0 all
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit // turn ON that bit 3 of word 1 (i.e. bit-67)
}

// convert set <s> to the "<s> UNION <t>"
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
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1,9,144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9,42}

	x.UnionWith(&y)
	fmt.Println(x.String()) // {1,9,42,144}

	fmt.Println(x.Has(9), y.Has(123)) // "true" "false"
}

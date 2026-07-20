package exercise112

import "testing"

// check if our custom "IntSet" matches a standard library "map[int]bool"
func equalSet(s *IntSet, m map[int]bool) bool {
	// check if each key in "map" exists in "s"
	for x := range m {
		if !s.Has(x) {
			return false
		}
	}
	// check if each element in "IntSet" exists in "map"
	for word, bits := range s.words {
		// word: row number
		// bits: 64 bits
		for bit := 0; bit < 64; bit++ {
			// if a bit is ON
			if bits&(1<<uint(bit)) != 0 {
				x := word*64 + bit
				if !m[x] {
					return false
				}
			}
		}
	}
	return true
}

func TestIntSet(t *testing.T) {
	var s IntSet
	m := make(map[int]bool)

	// Add elements to both
	values := []int{1, 9, 43, 64, 68}
	for _, x := range values {
		s.Add(x)
		m[x] = true
	}
	// test Add()
	if !equalSet(&s, m) {
		t.Fatalf("after Add()")
	}
	// test Has()
	for i := 0; i < 200; i++ {
		if s.Has(i) != m[i] {
			t.Fatalf("Has(%d): got %v want %v", i, s.Has(i), m[i])
		}
	}
	// test UnionWith()
	var tset IntSet
	m2 := make(map[int]bool)

	values2 := []int{9, 100, 101, 200}
	for _, x := range values2 {
		tset.Add(x)
		m2[x] = true
	}

	s.UnionWith(&tset)
	for x := range m2 {
		m[x] = true
	}

	if !equalSet(&s, m) {
		t.Fatal("after UnionWith()")
	}
}

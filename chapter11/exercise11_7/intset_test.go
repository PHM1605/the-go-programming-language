package exercise117

import (
	"math/rand"
	"testing"
)

// Generate random data sequence of 100000 numbers
var randomNumbers []int

func init() {
	randomNumbers = make([]int, 100000)
	// random generator
	r := rand.New(rand.NewSource(1))
	for i := range randomNumbers {
		randomNumbers[i] = r.Intn(1_000_000)
	}
}

func BenchmarkHas(b *testing.B) {
	// setup code
	var s IntSet
	for _, x := range randomNumbers {
		s.Add(x)
	}
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
		for _, x := range randomNumbers {
			s.Has(x)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s IntSet
		for _, x := range randomNumbers {
			s.Add(x)
		}
	}
}

func BenchmarkUnionWith(b *testing.B) {
	// setup code; create 2 sets
	var a, c IntSet
	for i, x := range randomNumbers {
		if i%2 == 0 {
			a.Add(x)
		} else {
			c.Add(x)
		}
	}
	b.ResetTimer()

	// benchmark
	for i := 0; i < b.N; i++ {
		var tmp IntSet
		tmp.words = append(tmp.words, a.words...)
		tmp.UnionWith(&c)
	}
}

// Compare against standard library
func BenchmarkMapHas(b *testing.B) {
	// setup code
	m := make(map[int]bool)
	for _, x := range randomNumbers {
		m[x] = true
	}
	b.ResetTimer()
	// benchmark
	for i := 0; i < b.N; i++ {
		for _, x := range randomNumbers {
			_, ok := m[x]
			_ = ok
		}
	}
}
func BenchmarkMapAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]bool)
		for _, x := range randomNumbers {
			m[x] = true
		}
	}
}

package exercise115

import (
	"reflect"
	"strings"
	"testing"
)

// // Bad way
// func assertEqual(x, y int) {
// 	if x != y {
// 		panic(fmt.Sprintf("%d != %d", x, y))
// 	}
// }

// func TestSplit(t *testing.T) {
// 	words := strings.Split("a:b:c", ":")
// 	assertEqual(len(words), 4)
// }

func TestSplitSingle(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	// make it FAIL on purpose
	if got, want := len(words), 4; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
	}
}

func TestSplitMultiple(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"abc", ",", []string{"abc"}},
		{"", ",", []string{""}},
		{"a::b", ":", []string{"a", "", "b"}},
	}

	for _, test := range tests {
		got := strings.Split(test.s, test.sep)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Split(%q, %q) = %v, want %v", test.s, test.sep, got, test.want)
		}
	}
}

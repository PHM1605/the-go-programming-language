package charcount

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	// UTFMax: utf8 has maximum 4 bytes => we need 5 ints
	// utflen: count how many runes have 1,2,3,4 bytes (0 is unused)
	counts, utflen, invalid, err := CharCount(bytes.NewBufferString("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if invalid != 0 {
		t.Fatalf("invalid=%d want=0", invalid)
	}
	if counts['h'] != 1 {
		t.Errorf("h=%d want=1", counts['h'])
	}
	if counts['e'] != 1 {
		t.Errorf("e=%d want=1", counts['e'])
	}
	if counts['l'] != 2 {
		t.Errorf("l=%d want=2", counts['l'])
	}
	if counts['o'] != 1 {
		t.Errorf("o=%d want=1", counts['o'])
	}
	// number of char of size=1 is 5 (of "hello" string)
	if utflen[1] != 5 {
		t.Errorf("utflen[1]=%d want=5", utflen[1])
	}
}

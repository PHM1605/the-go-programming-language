package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

func CharCount(r io.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int, error) {
	counts := make(map[rune]int)
	invalid := 0 // how many invalid runes
	// UTFMax: utf8 has maximum 4 bytes => we need 5 ints
	// utflen: count how many runes have 1,2,3,4 bytes (0 is unused)
	var utflen [utf8.UTFMax + 1]int

	in := bufio.NewReader(r)
	for {
		ch, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, utflen, invalid, err
		}
		if ch == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[ch]++
		utflen[n]++
	}
	return counts, utflen, invalid, nil
}

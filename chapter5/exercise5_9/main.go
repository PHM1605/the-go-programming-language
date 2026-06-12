package main

import (
	"fmt"
	"strings"
	"unicode"
)

// expand(): return a new long string
// f("bar") will decide what to do with string "$bar"
func expand(s string, f func(string) string) string {
	var result strings.Builder
	for i := 0; i < len(s); {
		if s[i] != '$' {
			result.WriteByte(s[i])
			i++
			continue
		}
		// Found '$' => try finding key e.g. $foo to be passed to f("foo")
		j := i + 1
		for j < len(s) && (unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j])) || s[j] == '_') {
			j++
		}
		// if $ is not followed by digit or character of '_'
		if j == i+1 {
			result.WriteByte('$')
			i++
			continue
		}

		keyword := s[i+1 : j]
		result.WriteString(f(keyword))
		i = j
	}
	return result.String()
}

func main() {
	s := "Hello $name, welcome to $language"
	out := expand(s, func(name string) string {
		switch name {
		case "name":
			return "Minh"
		case "language":
			return "Golang"
		default:
			return ""
		}
	})
	fmt.Println(out)
}

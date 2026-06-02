// Put "," in big numbers e.g. 1234567 => 1,234,567
// go run comma.go
package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		// no insertion of ','
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Printf("After insertion: %q\n", comma("1234567"))
}

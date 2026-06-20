package main

import "fmt"

type ByteCounter int

// to make it suitable for Fprintf()
// "how many bytes have been printed by Fprintf()"
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // to make 2 sides of the expression have same type "ByteCounter"
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // len("hello") = 5

	c = 0 // reset counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // len("hello, Dolly") = 12
}

// Car program to test "struct inside struct"
// go run embed.go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point  // NOTE: this field has only TYPE but NONAME
	Radius int
}

type Wheel struct {
	Circle     // NOTE: this field has only TYPE but NONAME
	Spokes int // number of rods in a wheel
}

func main() {
	// short form
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	// long form
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w1) // print in clear form i.e. with field names
	fmt.Printf("%v\n", w2)  // print in short form
	// => we can access DEEP INSIDE elements of w1 or w2
	fmt.Printf("%v\n", w1.X)
}

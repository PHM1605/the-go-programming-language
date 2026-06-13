package main

import (
	"fmt"
	. "geometry_project/geometry"
)

func main() {
	// Point distance
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(Distance(p, q)) // normal function
	fmt.Println(p.Distance(q))  // method

	// Perimeter calculation
	perim := Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println(perim.Distance())

	// Try "receiver is a pointer" (we call it "*Point method")
	r := &Point{X: 1, Y: 2}
	r.ScaleBy(2)
	fmt.Println(*r)
}

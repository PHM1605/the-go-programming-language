package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// METHODS
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	// Create a ColorPoint
	var cp ColoredPoint
	cp.X = 1
	// long and short form
	fmt.Println(cp.Point.X) // "1"
	fmt.Println(cp.X)       // "1"
	// long and short form
	cp.Point.Y = 2
	fmt.Println(cp.Point.Y) // "2"
	fmt.Println(cp.Y)       // "2"

	// Some colors tests
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{Point{X: 5, Y: 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}

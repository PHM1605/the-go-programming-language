package geometry

import "math"

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method to Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// another type, list of many Points, that has "Distance" too
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Method with POINTER to a type
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

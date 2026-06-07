// Modify "corner()" in "surface" in Chapter 3 to use "named results" and "bare return" statement
// go run main.go > surface.svg
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // xy = 600 pixels, z = 320 pixels
	cells         = 100      // #grid-cells; 100 rows 100 cols
	xyrange       = 30.0     // axis range (-xyrange/2..+xyrange/2)
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // distance from that point to origin
	return math.Sin(r) / r // sin(distance)/distance
}

// NEW: "named result" here
func corner(i, j int) (sx, sy float64) {
	// -0.5->0.5 of xyrange = -0.5->0.5 of 30 = -15->15
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	// from 3D to 2D (sx,sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	// NEW: "bare return" statement here
	return
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill:white; stroke-width:0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

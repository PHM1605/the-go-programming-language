package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100  // 100 cell rows, 100 cell columns
	xyrange       = 30.0 // xy from [-15..15]
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// convert 4 corners of cell having index of (i,j) to its location on 2D
func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	// xyz in 3D
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	// xy in 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(out io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill: white; stroke-width:0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g, %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

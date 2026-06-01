// go run exercise3_2.go
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 20.0                // both x and y are in range -15..15
	xyscale       = width / 2 / xyrange // xscale = 10, yscale = 5.333
	zscale        = height * 0.4        // zcale = 128
	angle         = math.Pi / 6         // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// NEW: corners calculation takes z=f(x,y) as a parameter
func corner(f func(float64, float64) float64, i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5) // [-15..15]
	y := xyrange * (float64(j)/cells - 0.5) // [-15..15]
	z := f(x, y)
	// from 3D to 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// generate .svg image from a corner function
func generate(filename string, f func(float64, float64) float64) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(
		file,
		"<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n",
		width,
		height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)

			fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(file, "</svg>")
}

func main() {
	// Egg box shape
	eggbox := func(x, y float64) float64 {
		return math.Sin(x) * math.Sin(y)
	}
	// Moguls shape
	moguls := func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r)
	}
	// Shaddle shape
	saddle := func(x, y float64) float64 {
		return (x*x - y*y) / 50
	}
	generate("surface1.svg", eggbox)
	generate("surface2.svg", moguls)
	generate("surface3.svg", saddle)
}

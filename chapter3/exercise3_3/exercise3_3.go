// Improve: provide COLOR CODE for Z-value
// go run exercise3_3.go > surface.svg
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // xy = 600 pixels, z = 320 pixels
	cells         = 100                 // number of grid cells; 100 rows 100 columns
	xyrange       = 30.0                // axis ranges (-xyrange/2..+xyrange/2)
	xyscale       = width / 2 / xyrange // -30..30 == -300pixels..300pixels
	zscale        = height * 0.4        // zscale = 128 => height = 2.5 units
	angle         = math.Pi / 6         // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// NEW: add HEIGHT as well for color code
func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5) // -0.5->0.5 of xyrange i.e. -15->15
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// from 3D (x,y,z) to 2D (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // distance from that point to origin
	return math.Sin(r) / r // sin(distance) / distance
}

// NEW: color mapping function
func color(z float64) string {
	// z ranges [-1,1] => normalize to [0,1]
	norm := (z + 1) / 2
	if norm < 0 {
		norm = 0
	}
	if norm > 1 {
		norm = 1
	}
	r := int(255 * norm)
	b := int(255 * (1 - norm))
	// RGB
	return fmt.Sprintf("#%02x00%02x", r, b)
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			// average value of z of 4 corners
			zavg := (az + bz + cz + dz) / 4
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color(zavg))
		}
	}
	fmt.Println("</svg>")
}

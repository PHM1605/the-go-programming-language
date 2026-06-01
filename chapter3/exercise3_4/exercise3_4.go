// Construct a webserver that computes surfaces and writes SVG data to client
// usage: http://localhost:8000/?width=1000&height=6000&color=red
package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6 // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// surface function
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}

// project 3D → 2D
func corner(width, height int, xyscale, zscale float64, i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set response header to send .svg back
	w.Header().Set("Content-Type", "image/svg+xml")
	// use default values OR set parameters from request query
	width := 600
	height := 320
	color := "gray"
	if v := r.URL.Query().Get("width"); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			width = val
		}
	}
	if v := r.URL.Query().Get("height"); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			height = val
		}
	}
	if v := r.URL.Query().Get("color"); v != "" {
		color = v
	}
	// Calculate scales from width and height
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	// Outputing image
	fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' style='stroke: %s; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n",
		color, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(width, height, xyscale, zscale, i+1, j)
			bx, by := corner(width, height, xyscale, zscale, i, j)
			cx, cy := corner(width, height, xyscale, zscale, i, j+1)
			dx, dy := corner(width, height, xyscale, zscale, i+1, j+1)

			fmt.Fprintf(w,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}

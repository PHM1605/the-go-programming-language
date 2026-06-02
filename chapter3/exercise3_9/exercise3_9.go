// go run exercise3_9.go
// http://localhost:8000?x=-0.75&y=0.1&zoom=2
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
)

// z is in complex plane [-2..2] in both x and y
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func handler(w http.ResponseWriter, r *http.Request) {
	// to response that "we send an image back"
	w.Header().Set("Content-Type", "image/png")

	width, height := 1024, 1024
	centerX, centerY := 0.0, 0.0
	zoom := 1.0

	// read params from query
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
	if v := r.URL.Query().Get("x"); v != "" {
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			centerX = val
		}
	}
	if v := r.URL.Query().Get("y"); v != "" {
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			centerY = val
		}
	}
	if v := r.URL.Query().Get("zoom"); v != "" {
		if val, err := strconv.ParseFloat(v, 64); err == nil && val > 0 {
			zoom = val
		}
	}

	// compute viewing window
	baseRange := 2.0
	xmin := centerX - baseRange/zoom
	xmax := centerX + baseRange/zoom
	ymin := centerY - baseRange/zoom
	ymax := centerY + baseRange/zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// x, y: range [-2..2]
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			// set pixel (px,py) to complex value
			img.Set(px, py, mandelbrot(complex(x, y)))
		}
	}
	// output image to response writer
	png.Encode(w, img)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}

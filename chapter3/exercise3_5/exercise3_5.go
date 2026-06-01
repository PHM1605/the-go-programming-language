// color version of mandelbrot
// go run exercise3_5.go > mandelbrot.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	// side-by-side: left=RGBA, right=YCbCr
	img := image.NewRGBA(image.Rect(0, 0, width*2, height))
	// x, y: range [-2..2]
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// set pixel (px,py) to complex value
			img.Set(px, py, mandelbrotRGBA(z))
			img.Set(px+width, py, mandelbrotYCbCr(z))
		}
	}
	// output image to stdout
	png.Encode(os.Stdout, img)
}

func mandelbrotRGBA(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				R: 255 - 5*n,
				G: 10 * n,
				B: 128 + 2*n,
				A: 255,
			}
		}
	}
	return color.Black
}

func mandelbrotYCbCr(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{
				Y:  255 - 10*n,
				Cb: 128 + 5*n,
				Cr: 200 - 3*n,
			}
		}
	}
	return color.Black
}

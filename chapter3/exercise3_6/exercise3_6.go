// Supersampling is computing the color value at 4 points within each pixel and taking the average
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
	// left: no supersampling; right: supersampling
	img := image.NewRGBA(image.Rect(0, 0, width*2, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// left: no supersampling
			// x, y: range [-2..2]
			x := float64(px)/width*(xmax-xmin) + xmin
			y := float64(py)/height*(ymax-ymin) + ymin
			img.Set(px, py, mandelbrot(complex(x, y)))
			// right: supersamplingl using 4 SUBPIXELS
			// main pixel is (px, py)
			// 4 subpixels are (px+0.25,py+0.25), (px+0.75,py+0.25), (px+0.25,py+0.75), (px+0.75,py+0.75)
			offsets := [2]float64{0.25, 0.75}
			var totalR, totalG, totalB float64
			for dy := 0; dy < 2; dy++ {
				for dx := 0; dx < 2; dx++ {
					sx := (float64(px)+offsets[dx])/width*(xmax-xmin) + xmin
					sy := (float64(py)+offsets[dy])/height*(ymax-ymin) + ymin
					r, g, b, _ := mandelbrot(complex(sx, sy)).RGBA()
					totalR += float64(r)
					totalG += float64(g)
					totalB += float64(b)
				}
			}
			// average all channels; remember that RGBA() returns [0..65535] instead of [0..255]
			avgR := uint8(totalR / 4 / 256)
			avgG := uint8(totalG / 4 / 256)
			avgB := uint8(totalB / 4 / 256)
			img.Set(px+width, py, color.RGBA{avgR, avgG, avgB, 255})
		}
	}
	// export to image
	png.Encode(os.Stdout, img)
}

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

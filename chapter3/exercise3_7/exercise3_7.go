// Use Newton's method to find complex solutions to "z^4-1=0"
// Solution: Newton uses iteration: z = z-f(z)/f'(z); 4 roots are 1,-1,i,-i
// shade each point by NUMBER-OF-ITERATIONS to reach 1 of 4 roots
// run: go run exercise3_7.go > newton.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func newton(z complex128) color.Color {
	const iterations = 200
	const tolerance = 1e-6
	for n := uint8(0); n < iterations; n++ {
		z -= (cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3)) // z = z - f(z)/f'(z)
		// shade each point by number-of-iterations
		brightness := 255 - n
		// check which root of the 4 (1,-1,i,-i); color coding
		// the more bright the less iterations it needs
		switch {
		case cmplx.Abs(z-1) < tolerance:
			return color.RGBA{brightness, 0, 0, 255} // red: root z=1
		case cmplx.Abs(z+1) < tolerance:
			return color.RGBA{0, brightness, 0, 255} // green: root z=-1
		case cmplx.Abs(z-1i) < tolerance:
			return color.RGBA{0, 0, brightness, 255} // blue: root z=+i
		case cmplx.Abs(z+1i) < tolerance:
			return color.RGBA{brightness, brightness, 0, 255} // yellow: root z=-i
		}
	}
	return color.Black // didn't converge
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, newton(complex(x, y)))
		}
	}
	png.Encode(os.Stdout, img)
}

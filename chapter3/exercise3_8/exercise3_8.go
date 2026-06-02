// implement mandelbrot program using 4 different representations of numbers: complex64, complex128, big.Float, big.Rat
// go run exercise3_8.go > mandelbrot.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 400, 400
	iterations             = 200
	contrast               = 15
)

// complex64
func mandelbrot64(z complex64) color.Color {
	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		// squared of magnitude
		if real(v)*real(v)+imag(v)*imag(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// complex128
func mandelbrot128(z complex128) color.Color {
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// big.Float; using real part "xf" and imaginary part "yf"
func mandelbrotBigFloat(xf, yf float64, prec uint) color.Color {
	x := new(big.Float).SetPrec(prec).SetFloat64(xf) // SetFloat64() means "create from a float64"
	y := new(big.Float).SetPrec(prec).SetFloat64(yf)
	// vx and vy = 0 initially
	vx := new(big.Float).SetPrec(prec)
	vy := new(big.Float).SetPrec(prec)
	// some big float constants
	two := new(big.Float).SetPrec(prec).SetFloat64(2)
	four := new(big.Float).SetPrec(prec).SetFloat64(4)

	for n := uint8(0); n < iterations; n++ {
		// v*v+z with v=vx+j(vy)
		vx2 := new(big.Float).Mul(vx, vx)
		vy2 := new(big.Float).Mul(vy, vy)
		// => real is "vx^2-vy^2 + x"
		real := new(big.Float).Sub(vx2, vy2)
		real.Add(real, x)
		// imaginary is "2*vx*vy + y"
		imag := new(big.Float).Mul(vx, vy)
		imag.Mul(imag, two)
		imag.Add(imag, y)
		// new v
		vx, vy = real, imag
		// check condition to continue or return
		// mag = x^2+y^2
		mag := new(big.Float).Add(
			new(big.Float).Mul(vx, vx),
			new(big.Float).Mul(vy, vy),
		)
		if mag.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// big.Float; using real part "xf" and imaginary part "yf"
func mandelbrotBigRat(xf, yf float64) color.Color {
	x := new(big.Rat).SetFloat64(xf) // SetFloat64() means "create from a float64"
	y := new(big.Rat).SetFloat64(yf)
	// vx and vy = 0 initially
	vx := new(big.Rat)
	vy := new(big.Rat)
	// some big float constants
	two := big.NewRat(2, 1)  // 2/1=2
	four := big.NewRat(4, 1) // 4/1=4

	for n := uint8(0); n < iterations; n++ {
		// v*v+z with v=vx+j(vy)
		vx2 := new(big.Rat).Mul(vx, vx)
		vy2 := new(big.Rat).Mul(vy, vy)
		// => real is "vx^2-vy^2 + x"
		real := new(big.Rat).Sub(vx2, vy2)
		real.Add(real, x)
		// imaginary is "2*vx*vy + y"
		imag := new(big.Rat).Mul(vx, vy)
		imag.Mul(imag, two)
		imag.Add(imag, y)
		// new v
		vx, vy = real, imag
		// check condition to continue or return
		// mag = x^2+y^2
		mag := new(big.Rat).Add(
			new(big.Rat).Mul(vx, vx),
			new(big.Rat).Mul(vy, vy),
		)
		if mag.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {
	// 4 panels side-by-side
	img := image.NewRGBA(image.Rect(0, 0, width*4, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			y := float64(py)/height*(ymax-ymin) + ymin

			// complex64
			img.Set(px, py, mandelbrot64(complex64(complex(x, y))))
			// complex128
			img.Set(px+width, py, mandelbrot128(complex(x, y)))
			// big.Float
			img.Set(px+2*width, py, mandelbrotBigFloat(x, y, 128))
			// big.Rat
			img.Set(px+3*width, py, mandelbrotBigRat(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

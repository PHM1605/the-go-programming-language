package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 2048, 2048
)

// color value at a location "z" in the z-plane
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	// this value exceeding 2 will return color
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {
	// NEW: check #CPU-cores first
	numCPUs := runtime.NumCPU()
	fmt.Println("NumCPU: ", numCPUs)
	// NEW: check running with number of cores from 0 to GOMAXPROCS-1
	for p := 1; p <= numCPUs; p++ {
		runtime.GOMAXPROCS(p) // allocate number of CPU cores to use
		start := time.Now()
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// channel to distribute row-of-image (integer) to each goroutine
		rows := make(chan int)

		numGoroutines := 10
		// goroutine counter to make sure all finishes his job before exit
		var wg sync.WaitGroup

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			// start 10 goroutines
			go func() {
				defer wg.Done()
				// get row number (from 0..2048 randomly) from work-channel and process
				// 1 goroutine will get 1 row-number
				for py := range rows {
					// convert from pixel value to numeric
					y := float64(py)/height*(ymax-ymin) + ymin
					for px := 0; px < width; px++ {
						// convert from pixel value to numeric
						x := float64(px)/width*(xmax-xmin) + xmin
						z := complex(x, y)
						img.Set(px, py, mandelbrot(z))
					}
				}
			}()
		}

		// distribute 1 row number to 1 goroutine
		for py := 0; py < height; py++ {
			rows <- py
		}
		// wait for all goroutines to finish their jobs
		close(rows)
		wg.Wait()

		// conclude
		png.Encode(os.Stdout, img)
		fmt.Fprintf(os.Stderr, "GOMAXPROCS=%d time=%v\n", p, time.Since(start))
	}
}

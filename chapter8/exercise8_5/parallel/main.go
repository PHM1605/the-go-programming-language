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
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// NOTE: channel for 10 goroutines to receive ROW NUMBER to care for
	// e.g. goroutine 7 receives row number 278, then 576... => calculate mandelbrot value for that entire row
	rows := make(chan int)

	// we use as many goroutines as our Machine can provide => 10 workers in my Mac
	numWorkers := runtime.NumCPU()
	// to control that all 10 workers have finished their jobs
	var wg sync.WaitGroup
	// create 10 workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			// mark Done at the end to countdown WaitGroup
			defer wg.Done()
			// actual work of calculate Mandelbrot value for row "py"
			// NOTE: this will NOT run once; but keep looping until "row" channel is closed
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}

	// distribute row number to channel (to be shared by 10 goroutines)
	for py := 0; py < height; py++ {
		rows <- py
	}

	// wait for all goroutines to finish their jobs
	close(rows)
	wg.Wait()

	// conclude
	fmt.Fprintf(os.Stderr, "Parallel time: %v\n", time.Since(start))
	png.Encode(os.Stdout, img)
}

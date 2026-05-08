// Server returns a Lissajous every time visiting "/"
// $ go run exercise1_12.go
// Note: visiting http://localhost:8000/?cycles=20 set cycles to 20 instead of default 5
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // black
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
}

const (
	blackIndex = 0 // 1st color in palette
	greenIndex = 1 // 2nd color in palette
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Take cycle parameter from request Query
		cycles := 5
		if c := r.URL.Query().Get("cycles"); c != "" {
			if n, err := strconv.Atoi(c); err == nil {
				cycles = n
			}
		}

		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100 // image canvas size [-size..size] for both axes
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // gif has 64 frames
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette) // palette default background of 0 => black
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

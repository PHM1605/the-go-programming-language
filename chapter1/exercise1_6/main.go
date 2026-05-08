// GIF of Lissajous curves
// NEW: change to multiple colors
// $ go build main.go
// $ ./main >out.gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // black
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // cyan
}

const (
	blackIndex = 0 // 1st color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // #cycles
		res     = 0.001 // resolution
		size    = 100   // image canvas has size [-size..size] for both axis
		nframes = 64    // #animation frames
		delay   = 8     // period between 2 frames; in [10ms] unit
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // gif has 64 frames
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette) // palette default background of 0 => color black index
		// NEW: choose color
		colorIndex := uint8((i % (len(palette) - 1)) + 1) // if 7 colors => choose index 1->6 (not choose background color of black)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

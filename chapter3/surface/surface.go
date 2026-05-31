package main 

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // xy = 600 pixels, z = 320 pixels
	cells = 100 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // -30..30 == -300pixels..300pixels
	zscale = height*0.4 // zscale = 128 => height = 2.5 units
	angle = math.Pi / 6 // 30 degrees
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("")
}
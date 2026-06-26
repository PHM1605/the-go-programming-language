package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// producing smaller thumbnail image from big image
func ImageFile(infile string) (string, error) {
	// open original image
	f, err := os.Open(infile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// decode JPEG into binary
	src, err := jpeg.Decode(f)
	if err != nil {
		return "", err
	}
	// resize image
	thumb := resize(src, 128)
	// create output filename
	ext := filepath.Ext(infile)
	base := strings.TrimSuffix(infile, ext)
	outfile := base + ".thumb" + ext

	// create output file
	out, err := os.Create(outfile)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// encode thumbnail as JPEG
	if err := jpeg.Encode(out, thumb, nil); err != nil {
		return "", err
	}
	return outfile, nil
}

func resize(src image.Image, maxDim int) image.Image {
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	// calculate destination width and height (based on maximum limit of either)
	var dstW, dstH int
	if srcW > srcH { // horizontal image
		dstW = maxDim
		dstH = srcH * maxDim / srcW
	} else {
		dstH = maxDim
		dstW = srcW * maxDim / srcH
	}
	if dstW < 1 {
		dstW = 1
	}
	if dstH < 1 {
		dstH = 1
	}

	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	for y := 0; y < dstH; y++ {
		for x := 0; x < dstW; x++ {
			// which point in Source to choose
			srcX := bounds.Min.X + x*srcW/dstW
			srcY := bounds.Min.Y + y*srcH/dstH
			dst.Set(x, y, src.At(srcX, srcY))
		}
	}
	return dst
}

// make thumbnails in parallel for EACH file received from CHANNEL
// NEW: returns #bytes occupied by the files it creates
func makeThumbnails6(filenames <-chan string) int64 {
	// NEW: number of working goroutines
	var wg sync.WaitGroup
	// save sizes of processed files
	sizes := make(chan int64)

	// goroutines to process images
	for f := range filenames {
		wg.Add(1) // increment counter
		go func(f string) {
			defer wg.Done() // decrement counter
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	// goroutine for closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	// add up total number of bytes
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	dir := "images/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// NEW: use a separate routine to dump filename
	filenames := make(chan string)
	go func() {
		defer close(filenames)
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			// check if jpeg or jpg files
			ext := strings.ToLower(filepath.Ext(e.Name()))
			if ext != ".jpg" && ext != ".jpeg" {
				continue
			}
			// if that's image thumbnail => skip
			if strings.Contains(e.Name(), ".thumb.") {
				continue
			}
			// hang here until that filename is received (<-chan)
			filenames <- filepath.Join(dir, e.Name())
		}
	}()

	total := makeThumbnails6(filenames)
	fmt.Printf("Total thumbnail size: %d bytes\n", total)
}

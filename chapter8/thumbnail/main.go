package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
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

func makeThumbnails(filenames []string) {
	// make a shared channel among ALL goroutines to signal its completion
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)     // create thumbnail ~1s
			ch <- struct{}{} // signal "DONE" for this file
		}(f)
	}
	// wait for all goroutines to complete
	for range filenames {
		<-ch // this "receive" on an unbuffered channel will block until something is sent for EACH file
	}
}

// make thumbnails in parallel
// NEW: use a STRUCT as completion signal for each goroutine
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	// NEW: use buffered channel for ALL files
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it // signalling done for this goroutine
		}(f)
	}
	// waiting to collect responses from all goroutines
	for range filenames {
		// collecting DONE signal from all goroutines
		// when receive e.g. 7 over 10 needed => block when ch is empty
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func main() {
	dir := "images/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var images []string
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
		images = append(images, filepath.Join(dir, e.Name()))
	}

	// makeThumbnails(images)

	makeThumbnails5(images)
}

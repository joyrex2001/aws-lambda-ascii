package asciiart

/*
asciiart is a small library that can convert image.Image objects to ascii art.
It is heavily based upon: https://github.com/stdupp/goasciiart ()
*/

import (
	"github.com/nfnt/resize"

	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

// scale will rescale the image to provided width.
func scale(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

// convert will convert the image to ascii art, given a width, height and
// ascii art gradient string.
func convert(img image.Image, w, h int, gradient string) []string {
	out := []string{}
	for i := 0; i < h; i++ {
		line := ""
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := g.(color.Gray).Y
			pos := int(float32(y) / float32(len(gradient)))
			line = line + gradient[pos:pos+1]
		}
		out = append(out, line)
	}
	return out
}

// Convert will convert the given image to an ascii art image with given
// width and gradient string.
func Convert(img image.Image, width int, gradient string) []string {
	si, w, h := scale(img, width)
	return convert(si, w, h, gradient)
}

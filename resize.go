package imageutils

import (
	"image"
	"image/color"
)

type resized struct {
	src   image.Image
	scale int
}

// ColorModel implements image.Image interface
func (r resized) ColorModel() color.Model {
	return r.src.ColorModel()
}

// Bounds implements image.Image interface
func (r resized) Bounds() image.Rectangle {
	b := r.src.Bounds()
	return image.Rectangle{
		Min: image.Point{
			X: b.Min.X,
			Y: b.Min.Y,
		},
		Max: image.Point{
			X: b.Max.X * r.scale,
			Y: b.Max.Y * r.scale,
		},
	}
}

// At implements image.Image interface
func (r resized) At(x, y int) color.Color {
	return r.src.At(x/r.scale, y/r.scale)
}

// Resize resizes image to a given scale
// and then returns image.
// For now it will work only with positive integers.
func Resize(img image.Image, scale int) image.Image {
	if scale < 1 {
		scale = 1
	}
	return resized{img, scale}	
}

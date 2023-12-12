package imageutils

import (
	"image"
	"image/color"
)

type rescaled struct {
	src   image.Image
	scale int
}

// ColorModel implements image.Image interface
func (r rescaled) ColorModel() color.Model {
	return r.src.ColorModel()
}

// Bounds implements image.Image interface
func (r rescaled) Bounds() image.Rectangle {
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
func (r rescaled) At(x, y int) color.Color {
	return r.src.At(x/r.scale, y/r.scale)
}

// Scale scales image.Image to a given scale
// and then returns image.Image.
// For now it will work only with positive integers.
func Scale(img image.Image, scale int) image.Image {
	if scale < 1 {
		scale = 1
	}
	return rescaled{img, scale}
}

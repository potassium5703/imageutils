//go:build todo
package imageutils

import (
	"image"
	"image/color"
	"image/draw"
)

type pair struct {
	first, second image.Image
}

func (p pair) ColorModel() color.Model { return p.first.ColorModel() }

func (p pair) Bounds() image.Rectangle {
	var (
		b1 = p.first.Bounds()
		b2 = p.second.Bounds()
	)

	return image.Rectangle{
		image.ZP,
		image.Point{
			X: b1.Dx() + b2.Dx(),
			Y: max(b1.Dy(), b2.Dy()),
		},
	}
}

func (p pair) At(x, y int) color.Color {
	rgbaImg := image.NewRGBA(p.Bounds())
	draw.Draw(
		rgbaImg,
		rgbaImg.Bounds(),
		p.first,
		image.ZP,
		draw.Src,
	)
	draw.Draw(
		rgbaImg,
		p.Bounds(),
		p.second,
		image.Point{
			p.first.Bounds().Dx(),
			0,
		},
		draw.Src,
	)
	return rgbaImg
}

// Concat concatenates second image to the right side of first.
func Concat(i1, i2 image.Image) image.Image {
	return pair{i1, i2}
}

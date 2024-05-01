package imageutils

import (
	"image"
	"image/color"
	"image/draw"
)

type Side int

const (
	Left Side = iota
	Right 
	Up
	Down
)

type pair struct {
	first, second image.Image
	side          Side
}

// Takes color.Model of the first Image.
func (p pair) ColorModel() color.Model { return p.first.ColorModel() }

func (p pair) Bounds() image.Rectangle {
	var (
		b1 = p.first.Bounds()
		b2 = p.second.Bounds()
	)

	rect := image.Rectangle{}
	switch p.side {
	case Left:
		fallthrough
	case Right:
		rect = image.Rectangle{
			image.ZP,
			image.Point{
				X: b1.Dx() + b2.Dx(),
				Y: max(b1.Dy(), b2.Dy()),
			},
		}
	case Up:
		fallthrough
	case Down:
		rect = image.Rectangle{
			image.ZP,
			image.Point{
				X: max(b1.Dx(), b2.Dx()),
				Y: b1.Dy() + b2.Dy(),
			},
		}
	}
	return rect
}

func (p pair) At(x, y int) color.Color {
	rgbaImg := image.NewRGBA(p.Bounds())

	switch p.side {
	case Left:
		draw.Draw(
			rgbaImg, rgbaImg.Bounds(),
			p.second,
			image.ZP, draw.Src,
		)
		draw.Draw(
			rgbaImg, p.Bounds(),
			p.first,
			image.Point{
				p.second.Bounds().Dx(),
				0,
			},
			draw.Src,
		)
	case Right:
		draw.Draw(
			rgbaImg, rgbaImg.Bounds(),
			p.first,
			image.ZP, draw.Src,
		)
		draw.Draw(
			rgbaImg, p.Bounds(),
			p.second,
			image.Point{
				p.first.Bounds().Dx(),
				0,
			},
			draw.Src,
		)
	case Up:
		draw.Draw(
			rgbaImg, rgbaImg.Bounds(),
			p.first,
			image.ZP, draw.Src,
		)
		draw.Draw(
			rgbaImg, p.Bounds(),
			p.second,
			image.Point{
				0,
				p.first.Bounds().Dy(),
			},
			draw.Src,
		)
	case Down:
		draw.Draw(
			rgbaImg, rgbaImg.Bounds(),
			p.second,
			image.ZP, draw.Src,
		)
		draw.Draw(
			rgbaImg, p.Bounds(),
			p.first,
			image.Point{
				0,
				p.second.Bounds().Dy(),
			},
			draw.Src,
		)
	}
	return rgbaImg.At(x, y)
}

// Concat concatenates second image on a given side.
func Concat(i1, i2 image.Image, s Side) image.Image {
	return pair{i1, i2, s}
}

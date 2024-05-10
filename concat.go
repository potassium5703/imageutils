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

	point := image.Point{}
	switch p.side {
	case Left:
		fallthrough
	case Right:
		point = image.Point{
			X: b1.Dx() + b2.Dx(),
			Y: max(b1.Dy(), b2.Dy()),
		}
	case Up:
		fallthrough
	case Down:
		point = image.Point{
			X: max(b1.Dx(), b2.Dx()),
			Y: b1.Dy() + b2.Dy(),
		}
	}
	return image.Rectangle{
		image.ZP,
		point,
	}
}

func (p pair) At(x, y int) color.Color {
	img := image.NewRGBA(p.Bounds())

	point := image.Point{}
	switch p.side {
	case Left:
		p.first, p.second = p.second, p.first
		fallthrough
	case Right:
		point = image.Point{-p.first.Bounds().Dx(), 0}
	case Up:
		p.first, p.second = p.second, p.first
		fallthrough
	case Down:
		point = image.Point{0, -p.first.Bounds().Dy()}
	}
	// draw main
	draw.Draw(
		img, img.Bounds(),
		p.first,
		image.ZP, draw.Src,
	)

	// draw second
	draw.Draw(
		img, p.Bounds(),
		p.second,
		point,
		draw.Src,
	)

	return img.At(x, y)
}

func render(p pair) image.Image {
	img := image.NewRGBA(p.Bounds())

	point := image.Point{}
	switch p.side {
	case Left:
		p.first, p.second = p.second, p.first
		fallthrough
	case Right:
		point = image.Point{-p.first.Bounds().Dx(), 0}
	case Up:
		p.first, p.second = p.second, p.first
		fallthrough
	case Down:
		point = image.Point{0, -p.first.Bounds().Dy()}
	}
	// draw main
	draw.Draw(
		img, img.Bounds(),
		p.first,
		image.ZP, draw.Src,
	)

	// draw second
	draw.Draw(
		img, p.Bounds(),
		p.second,
		point,
		draw.Src,
	)

	return img
}

// Concat concatenates second image on a given side.
func Concat(i1, i2 image.Image, s Side) image.Image {
	return render(pair{i1, i2, s})
}

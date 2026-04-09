package graffiti

import (
	"image"
	"image/color"
)

// GradationImage returns a new gradation image of the given bounds and color.
func GradationImage(rectangle image.Rectangle, c color.Color) image.Image {
	return image.Image(gradationRGBA(rectangle, c))
}

func gradationRGBA(rectangle image.Rectangle, c color.Color) *image.RGBA {
	r, g, b, a := c.RGBA()

	var rWhite uint8 = 0xff
	distR := rWhite - uint8(r)
	distG := rWhite - uint8(g)
	distB := rWhite - uint8(b)
	a8 := uint8(a)

	rgba := image.NewRGBA(rectangle)
	dx := rectangle.Dx()
	dy := rectangle.Dy()
	stride := rgba.Stride

	// build first row
	for x := 0; x < dx; x++ {
		rate := float64(x) / float64(dx)
		off := x * 4
		rgba.Pix[off] = rWhite - uint8(float64(distR)*rate)
		rgba.Pix[off+1] = rWhite - uint8(float64(distG)*rate)
		rgba.Pix[off+2] = rWhite - uint8(float64(distB)*rate)
		rgba.Pix[off+3] = a8
	}

	// copy first row to remaining rows
	firstRow := rgba.Pix[:stride]
	for y := 1; y < dy; y++ {
		copy(rgba.Pix[y*stride:y*stride+stride], firstRow)
	}

	return rgba
}

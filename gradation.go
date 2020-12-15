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
	minX := rectangle.Min.X

	for x := minX; x < rectangle.Max.X; x++ {
		rate := float64(x-minX) / float64(dx)
		r := rWhite - uint8(float64(distR)*rate)
		g := rWhite - uint8(float64(distG)*rate)
		b := rWhite - uint8(float64(distB)*rate)
		for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
			rgba.Set(x, y, color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a8,
			})
		}
	}

	return rgba
}

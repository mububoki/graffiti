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

	rWhite64 := float64(0xffff)
	r64 := float64(r)
	g64 := float64(g)
	b64 := float64(b)
	a8 := uint8(a)

	rgba := image.NewRGBA(rectangle)
	dx := rectangle.Dx()
	minX := rectangle.Min.X

	for x := minX; x < rectangle.Max.X; x++ {
		rate := float64(x-minX) / float64(dx)
		r := uint8(rWhite64*(1-rate) + r64*rate)
		g := uint8(rWhite64*(1-rate) + g64*rate)
		b := uint8(rWhite64*(1-rate) + b64*rate)
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

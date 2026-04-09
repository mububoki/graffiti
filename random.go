package graffiti

import (
	"image"
	"image/color"
	"math/rand"
)

//RandomImage returns a new random Image.
func RandomImage(rectangle image.Rectangle) image.Image {
	rgba := image.NewRGBA(rectangle)
	pix := rgba.Pix
	for i := 0; i < len(pix); i += 4 {
		v := rand.Uint32()
		pix[i] = uint8(v)
		pix[i+1] = uint8(v >> 8)
		pix[i+2] = uint8(v >> 16)
		pix[i+3] = 0xff
	}
	return rgba
}

//RandomColor returns new random color
func RandomColor() color.Color {
	v := rand.Uint32()
	return &color.RGBA{
		R: uint8(v),
		G: uint8(v >> 8),
		B: uint8(v >> 16),
		A: 0xff,
	}
}

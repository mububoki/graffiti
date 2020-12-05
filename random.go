package graffiti

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomImage returns a new random Image.
func RandomImage(rectangle image.Rectangle) image.Image {
	return image.Image(randomImageRGBA(rectangle))
}

func randomImageRGBA(r image.Rectangle) *image.RGBA {
	rgba := image.NewRGBA(r)
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			rgba.Set(x, y, RandomColor())
		}
	}

	return rgba
}

//RandomColor returns new random color
func RandomColor() color.Color {
	return color.Color(randomColorRGBA())
}

func randomColorRGBA() *color.RGBA {
	return &color.RGBA{
		R: randomUint8(),
		G: randomUint8(),
		B: randomUint8(),
		A: randomUint8(),
	}
}

func randomUint8() uint8 {
	return uint8(rand.Uint32() % (1 << 8))
}

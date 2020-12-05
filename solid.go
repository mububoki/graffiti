package graffiti

import (
	"image"
	"image/color"
)

//NewUniform returns a new Solid image of the given bounds and color.
func SolidImage(rectangle image.Rectangle, c color.Color) image.Image {
	return image.Image(image.NewPaletted(rectangle, []color.Color{c}))
}

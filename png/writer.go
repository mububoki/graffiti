package png

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/mububoki/graffiti"
)

// EncodeSolid writes the solid image with the rectangle and the color c to w in PNG format.
func EncodeSolid(w io.Writer, rectangle image.Rectangle, c color.Color) error {
	return png.Encode(w, graffiti.SolidImage(rectangle, c))
}

// EncodeRandom writes the random image with the rectangle to w in PNG format.
func EncodeRandom(w io.Writer, rectangle image.Rectangle) error {
	return png.Encode(w, graffiti.RandomImage(rectangle))
}

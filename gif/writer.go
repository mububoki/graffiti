package gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	"github.com/mububoki/graffiti"
)

type Options = gif.Options

//EncodeSolid writes the solid colored image with the rectangle and the color c to w in GIF format.
func EncodeSolid(w io.Writer, rectangle image.Rectangle, c color.Color, o *Options) error {
	//plan9 := color.Palette(palette.Plan9)
	return gif.Encode(w, graffiti.SolidImage(rectangle, c), o)
}

//EncodeRandom writes the random image with the rectangle to w in GIF format.
func EncodeRandom(w io.Writer, rectangle image.Rectangle, o *Options) error {
	return gif.Encode(w, graffiti.RandomImage(rectangle), o)
}

package jpeg

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"

	"github.com/mububoki/graffiti"
)

type Options = jpeg.Options

//EncodeSolid writes the solid colored image with the rectangle and the color c to w in JPEG 4:2:0 baseline format with the given options. Default parameters are used if a nil *Options is passed.
func EncodeSolid(w io.Writer, rectangle image.Rectangle, c color.Color, o *Options) error {
	return jpeg.Encode(w, graffiti.SolidImage(rectangle, c), o)
}

//EncodeRandom writes the random image with the rectangle to w in JPEG 4:2:0 baseline format with the given options. Default parameters are used if a nil *Options is passed.
func EncodeRandom(w io.Writer, rectangle image.Rectangle, o *Options) error {
	return jpeg.Encode(w, graffiti.RandomImage(rectangle), o)
}

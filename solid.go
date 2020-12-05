package graffiti

import (
	"image"
	"image/color"
)

//Solid is a finite-sized Image of uniform color. It implements the color.Color, color.Model, and Image interfaces.
type Solid struct {
	*image.Uniform
	rectangle image.Rectangle
}

func (s *Solid) Bounds() image.Rectangle {
	return s.rectangle
}

//NewUniform returns a new Solid image of the given bounds and color.
func NewSolid(rectangle image.Rectangle, c color.Color) *Solid {
	return &Solid{
		Uniform:   image.NewUniform(c),
		rectangle: rectangle,
	}
}

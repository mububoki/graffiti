package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/mububoki/graffiti"
)

func main() {
	width, height := 10, 10
	rectangle := image.Rect(0, 0, width, height)

	randomImage := graffiti.RandomImage(rectangle)
	solidImage := graffiti.NewSolid(rectangle, color.White)

	fmt.Println(randomImage, solidImage)
}

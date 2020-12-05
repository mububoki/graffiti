package main

import (
	"fmt"
	"image"

	"github.com/mububoki/graffiti"
)

func main() {
	width, height := 1280, 720
	rectangle := image.Rect(0, 0, width, height)

	randomImage := graffiti.RandomImage(rectangle)
	solidImage := graffiti.SolidImage(rectangle, graffiti.RandomColor())

	fmt.Println(randomImage, solidImage)
}

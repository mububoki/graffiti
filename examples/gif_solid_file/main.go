package main

import (
	"image"
	"os"

	"github.com/mububoki/graffiti"
	"github.com/mububoki/graffiti/gif"
)

func main() {
	fileRandom, err := os.Create("solid.gif")
	if err != nil {
		panic(err)
	}
	defer fileRandom.Close()

	width, height := 1280, 720
	if err := gif.EncodeSolid(fileRandom, image.Rect(0, 0, width, height), graffiti.RandomColor(), &gif.Options{NumColors: 255}); err != nil {
		panic(err)
	}
}

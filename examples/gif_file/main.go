package main

import (
	"image"
	"os"

	"github.com/mububoki/graffiti"
	"github.com/mububoki/graffiti/gif"
)

func main() {
	fileRandom, err := os.Create("random.gif")
	if err != nil {
		panic(err)
	}
	defer fileRandom.Close()

	fileSolid, err := os.Create("solid.gif")
	if err != nil {
		panic(err)
	}
	defer fileSolid.Close()

	fileGradation, err := os.Create("gradation.gif")
	if err != nil {
		panic(err)
	}
	defer fileGradation.Close()

	width, height := 1280, 720
	if err := gif.EncodeRandom(fileRandom, image.Rect(0, 0, width, height), &gif.Options{NumColors: 255}); err != nil {
		panic(err)
	}

	if err := gif.EncodeSolid(fileSolid, image.Rect(0, 0, width, height), graffiti.RandomColor(), &gif.Options{NumColors: 255}); err != nil {
		panic(err)
	}

	if err := gif.EncodeGradation(fileGradation, image.Rect(0, 0, width, height), graffiti.RandomColor(), &gif.Options{NumColors: 255}); err != nil {
		panic(err)
	}
}

package main

import (
	"image"
	"os"

	"github.com/mububoki/graffiti"
	"github.com/mububoki/graffiti/png"
)

func main() {
	fileRandom, err := os.Create("random.png")
	if err != nil {
		panic(err)
	}
	defer fileRandom.Close()

	fileSolid, err := os.Create("solid.png")
	if err != nil {
		panic(err)
	}
	defer fileSolid.Close()

	fileGradation, err := os.Create("gradation.png")
	if err != nil {
		panic(err)
	}
	defer fileGradation.Close()

	width, height := 1280, 720
	if err := png.EncodeRandom(fileRandom, image.Rect(0, 0, width, height)); err != nil {
		panic(err)
	}

	if err := png.EncodeSolid(fileSolid, image.Rect(0, 0, width, height), graffiti.RandomColor()); err != nil {
		panic(err)
	}

	if err := png.EncodeGradation(fileGradation, image.Rect(0, 0, width, height), graffiti.RandomColor()); err != nil {
		panic(err)
	}
}

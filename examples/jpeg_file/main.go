package main

import (
	"image"
	"os"

	"github.com/mububoki/graffiti"
	"github.com/mububoki/graffiti/jpeg"
)

func main() {
	fileRandom, err := os.Create("random.jpg")
	if err != nil {
		panic(err)
	}
	defer fileRandom.Close()

	fileSolid, err := os.Create("solid.jpg")
	if err != nil {
		panic(err)
	}
	defer fileSolid.Close()

	fileGradation, err := os.Create("gradation.jpg")
	if err != nil {
		panic(err)
	}
	defer fileGradation.Close()

	width, height := 1280, 720
	if err := jpeg.EncodeRandom(fileRandom, image.Rect(0, 0, width, height), &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}

	if err := jpeg.EncodeSolid(fileSolid, image.Rect(0, 0, width, height), graffiti.RandomColor(), &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}

	if err := jpeg.EncodeGradation(fileGradation, image.Rect(0, 0, width, height), graffiti.RandomColor(), &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}
}

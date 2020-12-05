package main

import (
	"image"
	"os"

	"github.com/mububoki/graffiti/jpeg"
)

func main() {
	fileRandom, err := os.Create("random.jpg")
	if err != nil {
		panic(err)
	}
	defer fileRandom.Close()

	width, height := 1280, 720
	if err := jpeg.EncodeRandom(fileRandom, image.Rect(0, 0, width, height), &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}
}

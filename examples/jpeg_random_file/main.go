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

	if err := jpeg.EncodeRandom(fileRandom, image.Rect(0, 0, 100, 100), &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}
}

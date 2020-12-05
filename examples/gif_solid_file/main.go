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

	if err := gif.EncodeSolid(fileRandom, image.Rect(0, 0, 100, 100), graffiti.RandomColor(), nil); err != nil {
		panic(err)
	}
}

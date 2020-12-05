package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/mububoki/graffiti"
	"github.com/mububoki/graffiti/gif"
)

func main() {
	width, height := 1280, 720

	buf := new(bytes.Buffer)
	if err := gif.EncodeSolid(buf, image.Rect(0, 0, width, height), graffiti.RandomColor(), nil); err != nil {
		panic(err)
	}

	fmt.Println(buf.Bytes())
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

func ImageProfile(src io.Reader, path string, width, height int) {
	buf := make([]uint8, height*width)
	src.Read(buf)

	pic := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pic.Set(x, y, color.Gray{buf[y*width+x]})
		}
	}

	file, _ := os.Create(path)
	defer file.Close()
	png.Encode(file, pic)
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// Create a 100 by 50 pixels image
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	// Draw a red dot on img at pixel postion (2, 3)
	//
	// R - Red, G - Green, B - Blue, A - Alpha.
	//
	// In this case, 255 Red, 0 Green & Blue & 255 Alpha.
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})

	// Save img to a png file:

	// Create a file to save img into
	f, err := os.Create("img.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Encode and then write img to the file handle f
	png.Encode(f, img)
}

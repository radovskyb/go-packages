package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// Create a new image of size 128 by 128 pixels.
	img := image.NewRGBA(image.Rect(0, 0, 128, 128))

	// Encode and write img as a png to a file
	f, err := os.Create("gradient.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Draw a gradient on the blank image img
	drawGradient(img)

	// Encode m into png image format and store it in f
	if err := png.Encode(f, img); err != nil {
		log.Fatalln(err)
	}
}

func drawGradient(img *image.RGBA) {
	// Get the img's size (width and height)
	size := img.Bounds().Size()

	// Iterate over every pixel in img
	for x := 0; x < size.X; x++ { // Horizontal pixels
		for y := 0; y < size.Y; y++ { // Vertical pixels
			color := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255,
			}
			img.Set(x, y, color)
		}
	}
}

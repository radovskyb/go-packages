package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// Create a new rectangle image
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	// Create a colour blue to fill the image m with
	blue := color.RGBA{0, 0, 255, 255}

	// Draw the uniform colour blue into image rectangle m
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// Reset image m to transparent instead of blue, in this case
	// it will be uniform black since m's color model cannot represent
	// transparency.
	//
	// To reset an image to transparent (or black, if the destination
	// image's color model cannot represent transparency), use
	// image.Transparent, which is an image.Uniform:
	draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)

	f, err := os.Create("transparent_rectangle.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Encode the image to a jpeg file and write it to f
	if err := jpeg.Encode(f, m, nil); err != nil {
		log.Fatalln(err)
	}
}

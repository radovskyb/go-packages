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
	// NewRGBA returns a new RGBA image with the given bounds.
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	// RGBA represents a traditional 32-bit alpha-premultiplied color,
	// having 8 bits for each of red, green, blue and alpha.
	//
	// An alpha-premultiplied color component C has been scaled by
	// alpha (A), so has valid values 0 <= C <= A.
	blue := color.RGBA{0, 0, 255, 255}

	// Uniform is an infinite-sized Image of uniform color.
	// It implements the color.Color, color.Model, and Image interfaces.
	//
	// ZP is the zero Point.
	//
	// Src specifies ``src in mask''.
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// Create a new file to store the image m as a jpeg file
	f, err := os.Create("blue_square.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Encode the image m as a jpeg file and store it's contents in f
	if err := jpeg.Encode(f, m, nil); err != nil {
		log.Fatalln(err)
	}
}

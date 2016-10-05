package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	red   color.Color = color.RGBA{255, 0, 0, 255}
	blue  color.Color = color.RGBA{40, 100, 215, 255}
)

func main() {
	// Create a new image m
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	// Fill m in blue
	//
	// Draw calls DrawMask with a nil mask.
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// Draw a horizontal white line in m
	//
	// Set i to m's most minimal X's point `m.Bounds().Min.X` and loop
	// until i is at m's most maximum X's point, `m.Bounds().Max.X`, and
	// increment 1 each iteration and through each iteration, set the pixal
	// and the point at the position of i in image m to the colour white
	// at m's Y position at m's vertical half - 150 pixels.
	for i := m.Bounds().Min.X; i < m.Bounds().Max.X; i++ {
		m.Set(i, m.Bounds().Max.Y/2-150, white) // Set a single pixel
	}

	// Draw a vertical red dotted line in m
	for i := m.Bounds().Min.Y; i < m.Bounds().Max.Y; i += 3 {
		m.Set(m.Bounds().Max.X/2-225, i, red) // Set a single pixel
	}

	// Create a new file to store m in
	f, err := os.Create("lines.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Encode m into png image format and store it in f
	if err := png.Encode(f, m); err != nil {
		log.Fatalln(err)
	}
}

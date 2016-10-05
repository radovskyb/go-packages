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
	m1 := image.NewRGBA(image.Rect(0, 0, 640, 480))
	green := color.RGBA{200, 255, 25, 255}
	draw.Draw(m1, m1.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	m2 := image.NewRGBA(image.Rect(0, 0, 200, 200))
	blue := color.RGBA{0, 50, 255, 200}
	draw.Draw(m2, m2.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	m3 := m1.Bounds().Sub(m1.Rect.Min).Add(image.Point{50, 50})
	draw.Draw(m1, m3, m2, m1.Rect.Min, draw.Src)

	f, err := os.Create("m.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := jpeg.Encode(f, m1, nil); err != nil {
		log.Fatalln(err)
	}
}

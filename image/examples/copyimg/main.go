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
	src := image.NewRGBA(image.Rect(0, 0, 200, 400))
	blue := color.RGBA{0, 200, 255, 255}
	draw.Draw(src, src.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	f1, err := os.Create("f1.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	if err := jpeg.Encode(f1, src, nil); err != nil {
		log.Fatalln(err)
	}

	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	green := color.RGBA{150, 255, 25, 255}
	draw.Draw(dest, dest.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	f2, err := os.Create("f2.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f2.Close()

	if err := jpeg.Encode(f2, dest, nil); err != nil {
		log.Fatalln(err)
	}

	sr := src.Bounds()

	r := dest.Bounds().Sub(dest.Bounds().Min).Add(image.Point{50, 0})
	draw.Draw(dest, r, src, sr.Min, draw.Src)

	f3, err := os.Create("f3.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f3.Close()

	if err := jpeg.Encode(f3, dest, nil); err != nil {
		log.Fatalln(err)
	}
}

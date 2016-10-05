package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

// Circle will represent a circle.
type Circle struct {
	X, Y, R float64 // X - Width, Y - Height, R - Radius
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	}
	return uint8((1 - math.Pow(d, 5)) * 255) // Blur
}

func main() {
	var w, h int = 280, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	θ := 2 * math.Pi / 3

	// Create 3 new circles
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, c)
		}
	}

	f, err := os.Create("blured-circles.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := png.Encode(f, m); err != nil {
		log.Fatalln(err)
	}
}

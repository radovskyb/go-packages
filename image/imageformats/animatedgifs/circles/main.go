package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"os"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func main() {
	var w, h int = 240, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	circles := []*Circle{&Circle{}, &Circle{}, &Circle{}}

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	var images []*image.Paletted
	var delays []int
	steps := 20
	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		θ := 2.0 * math.Pi / float64(steps) * float64(step)
		for i, circle := range circles {
			θ0 := 2 * math.Pi / 3 * float64(i)
			circle.X = hw - 40*math.Sin(θ0) - 20*math.Sin(θ0+θ)
			circle.Y = hh - 40*math.Cos(θ0) - 20*math.Cos(θ0+θ)
			circle.R = 50
		}

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				img.Set(x, y, color.RGBA{
					circles[0].Brightness(float64(x), float64(y)),
					circles[1].Brightness(float64(x), float64(y)),
					circles[2].Brightness(float64(x), float64(y)),
					255,
				})
			}
		}
	}

	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	}); err != nil {
		log.Fatalln(err)
	}
}

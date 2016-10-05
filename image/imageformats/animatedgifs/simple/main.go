package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
)

func main() {
	// Create 3 colours
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create 3 new pallated images and store them in a slice
	images := []*image.Paletted{
		image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{red}),
		image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{green}),
		image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{blue}),
	}

	// Create a new &gif.GIF{} object
	outGif := &gif.GIF{}

	// Iterate over all of the pallated images in the images slice
	// and append each image and a delay for each image, to outGif
	for _, img := range images {
		// Append the image to the outGif object
		outGif.Image = append(outGif.Image, img)

		// Append a delay of 0 to the outGif object
		outGif.Delay = append(outGif.Delay, 25)
	}

	// Create a new file to store the outGif as a gif encoded file
	f, err := os.Create("images.gif")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Encode outGif as a gif file and store it in f
	if err := gif.EncodeAll(f, outGif); err != nil {
		log.Fatalln(err)
	}
}

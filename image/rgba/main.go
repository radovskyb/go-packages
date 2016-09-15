package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	// NewRGBA returns a new RGBA image with the given bounds.
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	m.Set(5, 5, color.RGBA{255, 0, 0, 255})

	m0 := image.NewRGBA(image.Rect(0, 0, 8, 5))

	// SubImage returns an image representing the portion of the image p visible
	// through r. The returned value shares pixels with the original image.
	m1 := m0.SubImage(image.Rect(1, 2, 5, 5)).(*image.RGBA)

	fmt.Println(m0.Bounds().Dx(), m1.Bounds().Dx()) // 8 4
	fmt.Println(m0.Stride == m1.Stride)             // true

	fmt.Println(m0.Bounds().Dy(), m1.Bounds().Dy()) // 5 3
}

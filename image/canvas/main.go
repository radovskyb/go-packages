package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type Vector struct {
	X, Y float64
}

func (v *Vector) Rotate(angle float64) {
	cos, sin := math.Cos(angle), math.Sin(angle)
	v.X, v.Y = v.X*cos+v.Y*sin, v.Y*cos-v.X*sin
}

func (v *Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y}
}

func (v *Vector) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v *Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vector) Scale(k float64) {
	v.X, v.Y = v.X*k, v.Y*k
}

func (v *Vector) toPoint() image.Point {
	return image.Point{int(v.X), int(v.Y)}
}

type Canvas struct {
	image.RGBA
}

// NewCanvas creates and initializes a new Canvas object with
// RGBA set to a new image.RGBA object that sets it's rect to r
func NewCanvas(r image.Rectangle) *Canvas {
	// Create a new canvas object
	canvas := new(Canvas)

	// Set canvas's RGBA object to a new image.RGBA object
	// that's rect is r
	canvas.RGBA = *image.NewRGBA(r)

	// Return the new canvas object
	return canvas
}

func (c *Canvas) DrawGradient() {
	size := c.Bounds().Size()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			colour := color.RGBA{
				uint8(255 * x / size.X),
				uint8(255 * y / size.Y),
				55,
				255,
			}
			c.Set(x, y, colour)
		}
	}
}

func (c *Canvas) DrawLine(colour color.RGBA, from Vector, to Vector) {
	delta := to.Sub(from)
	length := delta.Length()
	x_step, y_step := delta.X/length, delta.Y/length
	limit := int(length + 0.5)

	for i := 0; i < limit; i++ {
		x := from.X + float64(i)*x_step
		y := from.Y + float64(i)*y_step
		c.Set(int(x), int(y), colour)
	}
}

func (c *Canvas) DrawSpiral(colour color.RGBA, from Vector) {
	dir := Vector{0, 5}
	last := from

	for i := 0; i < 10000; i++ {
		next := last.Add(dir)
		c.DrawLine(colour, last, next)
		dir.Rotate(0.03)
		dir.Scale(0.999)
		last = next
	}
}

func main() {
	width, height := 125, 125

	// Draw simple gradient
	c := NewCanvas(image.Rect(0, 0, width, height))
	c.DrawGradient()

	f1, err := os.Create("canvas_gradient.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	if err := png.Encode(f1, c); err != nil {
		log.Fatalln(err)
	}

	// Draw lines
	c = NewCanvas(image.Rect(0, 0, width, height))
	c.DrawGradient()

	// Draw a series of lines from the top left corner to
	// the bottom of the image
	for x := 0; x < width; x += 8 {
		c.DrawLine(
			color.RGBA{0, 0, 0, 255},
			Vector{0.0, 0.0},
			Vector{float64(x), float64(height)},
		)
	}

	f2, err := os.Create("drawing_lines.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f2.Close()

	if err := png.Encode(f2, c); err != nil {
		log.Fatalln(err)
	}

	// Drawing spirals
	width, height = 2048, 1024

	c = NewCanvas(image.Rect(0, 0, width, height))
	c.DrawGradient()

	// Draw a set of spirals randomly over the image
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		x := float64(width) * rand.Float64()
		y := float64(height) * rand.Float64()

		colour := color.RGBA{
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			255,
		}

		c.DrawSpiral(colour, Vector{x, y})
	}

	f3, err := os.Create("spirals.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f3.Close()

	if err := png.Encode(f3, c); err != nil {
		log.Fatalln(err)
	}
}

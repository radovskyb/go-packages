package main

import (
	"fmt"
	"image"
)

func main() {
	// A Point is an X, Y coordinate pair. The axes increase right and down.
	p := image.Point{2, 1}
	fmt.Println(p)

	fmt.Println()

	// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned
	// rectangle has minimum and maximum coordinates swapped if necessary so that
	// it is well-formed.
	r := image.Rect(2, 1, 5, 5)

	// Dx and Dy return a rectangle's width and height.
	//
	// Dx returns r's width.
	//
	// Dy returns r's height.
	//
	// Pt is shorthand for Point{X, Y}.
	//
	// In reports whether p is in r.
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // 3 4 false

	fmt.Println()

	// Add returns the rectangle r translated by p.
	r = r.Add(image.Pt(-4, -2))
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // 3 4 true

	fmt.Println()

	// Intersecting two Rectangles yields another Rectangle, which may be empty.
	//
	// Intersect returns the largest rectangle contained by both r and s. If the
	// two rectangles do not overlap then the zero rectangle will be returned.
	r = image.Rect(0, 0, 4, 3).Intersect(image.Rect(2, 2, 5, 5))

	// Size returns a rectangle's width and height, as a Point.
	fmt.Printf("%#v\n", r.Size()) // prints image.Point{X:2, Y:1}
}

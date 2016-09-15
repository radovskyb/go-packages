package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var x int
	for i := 0; i < 1000000000; i++ {
		x++
	}
	fmt.Println(x)

	// Since returns the time elapsed since t.
	// It is shorthand for time.Now().Sub(t).
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}

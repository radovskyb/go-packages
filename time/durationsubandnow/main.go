package main

import (
	"fmt"
	"time"
)

func main() {
	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	fmt.Println(time.Duration(10) * time.Second)

	// Now returns the current local time.
	t0 := time.Now()

	var x int
	for i := 0; i < 1000000000; i++ {
		x++
	}
	fmt.Println(x)

	t1 := time.Now()

	// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
	// value that can be stored in a Duration, the maximum (or minimum) duration
	// will be returned.
	// To compute t-d for a duration d, use t.Add(-d).
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

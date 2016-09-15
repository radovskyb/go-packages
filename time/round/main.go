package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		// Round returns the result of rounding t to the nearest multiple
		// of d (since the zero time).
		// The rounding behavior for halfway values is to round up.
		// If d <= 0, Round returns t unchanged.
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
}

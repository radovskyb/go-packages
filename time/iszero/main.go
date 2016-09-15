package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time

	// See whether `t` has been set or still equals the zero time instant
	//
	// IsZero reports whether t represents the zero time instant,
	// January 1, year 1, 00:00:00 UTC.
	fmt.Println(t.IsZero()) // true

	// Add one hour to time `t`
	t = t.Add(time.Hour)

	fmt.Println(t.IsZero()) // false
}

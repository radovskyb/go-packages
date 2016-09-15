package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time and store it in `t0`
	t0 := time.Now()

	// Add 1 hour to the current time and store it in `t1`
	t1 := t0.Add(time.Hour)

	// Print whether t1's time is before t0's
	//
	// Before reports whether the time instant t is before u.
	fmt.Println(t1.Before(t0)) // false

	// Print whether t0's time is before t1's
	fmt.Println(t0.Before(t1)) // true

	fmt.Println()

	// Print whether t1's time is after t0's
	//
	// After reports whether the time instant t is after u.
	fmt.Println(t1.After(t0)) // true

	// Print whether t0's time is after t1's
	fmt.Println(t0.After(t1)) // false
}

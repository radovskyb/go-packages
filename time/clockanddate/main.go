package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time and store it in `t`
	t := time.Now()

	// Get the hours, minutes and seconds from `t`
	hours, minutes, seconds := t.Clock()

	// Print out the current hours, minutes and seconds returned from t.Clock()
	fmt.Printf("%v Hours, %v Minutes and %v Seconds.\n", hours, minutes, seconds)

	// Get the year, month and day from `t`
	//
	// Date returns the year, month, and day in which t occurs.
	year, month, day := t.Date()

	// Print out the current year, month and day returned from t.Clock()
	fmt.Printf("%v %v, %v.\n", month, day, year)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time and store it in `t`
	t := time.Now()

	// Weekday returns the day of the week specified by t.
	fmt.Println("Weekday:", t.Weekday())

	// YearDay returns the day of the year specified by t, in the
	// range [1,365] for non-leap years, and [1,366] in leap years.
	fmt.Println("Day of the year:", t.YearDay())
}

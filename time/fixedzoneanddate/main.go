package main

import (
	"fmt"
	"time"
)

func main() {
	// FixedZone returns a Location that always uses
	// the given zone name and offset (seconds east of UTC).
	gmt := time.FixedZone("GMT", 0)

	// String returns a descriptive name for the time zone information,
	// corresponding to the argument to LoadLocation.
	fmt.Println("Location:", gmt.String()) // Call String explicitly

	// Date returns the Time corresponding to
	// yyyy-mm-dd hh:mm:ss + nsec nanoseconds
	// in the appropriate zone for that time in the given location.
	//
	// The month, day, hour, min, sec, and nsec values may be outside
	// their usual ranges and will be normalized during the conversion.
	// For example, October 32 converts to November 1.
	//
	// A daylight savings time transition skips or repeats times.
	// For example, in the United States, March 13, 2011 2:15am never occurred,
	// while November 6, 2011 1:15am occurred twice.  In such cases, the
	// choice of time zone, and therefore the time, is not well-defined.
	// Date returns a time that is correct in one of the two zones involved
	// in the transition, but it does not guarantee which.
	//
	// Date panics if loc is nil.
	dayInGMT := time.Date(2015, 18, 5, 12, 15, 0, 0, gmt)
	fmt.Println(dayInGMT)

	fmt.Println()

	utc := time.FixedZone("UTC", 10*60*60)
	fmt.Println("Location:", utc)
	dayInUTC := time.Date(2015, 18, 5, 12, 15, 0, 0, utc)
	fmt.Println(dayInUTC)
}

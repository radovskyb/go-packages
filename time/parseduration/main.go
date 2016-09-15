package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	durationString := "75s"

	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of
	// decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatalln("ParseDuration error:", err)
	}

	// Print out the duration in hours
	//
	// Hours returns the duration as a floating point number of hours.
	fmt.Println(duration.Hours(), "hours")

	// Print out the duration in minutes
	//
	// Minutes returns the duration as a floating point number of minutes.
	fmt.Println(duration.Minutes(), "minutes")

	// Print out the duration in seconds
	//
	// Seconds returns the duration as a floating point number of seconds.
	fmt.Println(duration.Seconds(), "seconds")

	// Print out the duration in nanoseconds
	//
	// Nanoseconds returns the duration as an integer nanosecond count.
	fmt.Println(duration.Nanoseconds(), "nanoseconds")

	// Print out the duration as a string with the String() method explicitly
	//
	// String returns a string representing the duration in the form "72h3m0.5s".
	// Leading zero units are omitted.  As a special case, durations less than one
	// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
	// that the leading digit is non-zero.  The zero duration formats as 0,
	// with no unit.
	fmt.Println(duration.String())

	// Print out the duration as a string with the String() method implicitly
	fmt.Println(duration)
}

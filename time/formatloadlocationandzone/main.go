package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now:", now)

	fmt.Println()

	// Format returns a textual representation of the time value formatted
	// according to layout, which defines the format by showing how the reference
	// time, defined to be
	//	Mon Jan 2 15:04:05 -0700 MST 2006
	// would be displayed if it were the value; it serves as an example of the
	// desired output. The same display rules will then be applied to the time
	// value.
	//
	// A fractional second is represented by adding a period and zeros
	// to the end of the seconds section of layout string, as in "15:04:05.000"
	// to format a time stamp with millisecond precision.
	//
	// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
	// and convenient representations of the reference time. For more information
	// about the formats and the definition of the reference time, see the
	// documentation for ANSIC and the other constants defined by this package.
	fmt.Println(now.Format("3:04pm on Monday, January 2, 2006"))

	fmt.Println()

	// Location returns the time zone information associated with t.
	fmt.Println("Location:", now.Location())

	fmt.Println()

	// Get the time zone name and offset
	//
	// Zone computes the time zone in effect at time t, returning the abbreviated
	// name of the zone (such as "CET") and its offset in seconds east of UTC.
	zone, offset := now.Zone()

	fmt.Printf("Location (Time Zone): %v, Offset: %v\n", zone, offset)
}

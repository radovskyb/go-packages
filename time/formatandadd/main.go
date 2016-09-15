package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	t := time.Now()

	// Create a new time format to format time `t`
	format := "3:04pm"

	// Print the current time in the time format `format`
	//
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
	fmt.Println(t.Format(format))

	// Add 5 hours to time `t` and store it back into `t`
	//
	// Add returns the time t+d.
	t = t.Add(time.Hour * 5)

	// Print time `t` with the 5 hours added onto it
	fmt.Println(t.Format(format))
}

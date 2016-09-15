package main

import (
	"fmt"
	"log"
	"time"
)

const (
	longForm  = "Jan 2, 2006 at 3:04pm (MST)"
	shortForm = "2006-Jan-02"
)

func main() {
	// Parse parses a formatted string and returns the time value it represents.
	// The layout  defines the format by showing how the reference time,
	// defined to be
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// would be interpreted if it were the value; it serves as an example of
	// the input format. The same interpretation will then be made to the
	// input string.
	//
	// In the absence of a time zone indicator, Parse returns a time in UTC.
	//
	// When parsing a time with a zone offset like -0700, if the offset corresponds
	// to a time zone used by the current location (Local), then Parse uses that
	// location and zone in the returned time. Otherwise it records the time as
	// being in a fabricated location with time fixed at the given zone offset.
	//
	// No checking is done that the day of the month is within the month's
	// valid dates; any one- or two-digit value is accepted. For example
	// February 31 and even February 99 are valid dates, specifying dates
	// in March and May. This behavior is consistent with time.Date.
	//
	// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
	// has a defined offset in the current location, then that offset is used.
	// The zone abbreviation "UTC" is recognized as UTC regardless of location.
	// If the zone abbreviation is unknown, Parse records the time as being
	// in a fabricated location with the given zone abbreviation and a zero offset.
	// This choice means that such a time can be parsed and reformatted with the
	// same layout losslessly, but the exact instant used in the representation will
	// differ by the actual zone offset. To avoid such problems, prefer time layouts
	// that use a numeric zone offset, or use ParseInLocation.
	t, err := time.Parse(longForm, "Mar 16, 2016 at 3:55pm (AEDT)")
	if err != nil {
		log.Fatalln("Parse error:", err)
	}

	fmt.Println(t)
	t, err = time.Parse(shortForm, "2016-Mar-16")
	if err != nil {
		log.Fatalln("Parse error:", err)
	}
	fmt.Println(t)
}

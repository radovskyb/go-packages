package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// ISOWeek returns the ISO 8601 year and week number in which t occurs.
	// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
	// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
	// of year n+1.
	fmt.Println(t.ISOWeek())
}

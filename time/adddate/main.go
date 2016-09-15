package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current datetime and store it in variable `now`
	now := time.Now()

	// Print out `now`
	fmt.Println("Today:", now.Format(time.ANSIC))

	// Add 1 year to `now` and store the result in `addedDate`
	//
	// AddDate returns the time corresponding to adding the
	// given number of years, months, and days to t.
	// For example, AddDate(-1, 2, 3) applied to January 1, 2011
	// returns March 4, 2010.
	//
	// AddDate normalizes its result in the same way that Date does,
	// so, for example, adding one month to October 31 yields
	// December 1, the normalized form for November 31.
	addDate := now.AddDate(1, 0, 0)

	// Print out `addedDate` which will print 1 year after `now`'s datetime
	fmt.Println("After 1 year:", addDate.Format(time.ANSIC))
}

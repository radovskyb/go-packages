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
	loc, err := time.LoadLocation("Australia/Victoria")
	if err != nil {
		log.Fatalln("LoadLocation error:", err)
	}

	// ParseInLocation is like Parse but differs in two important ways.
	// First, in the absence of time zone information, Parse interprets a time as UTC;
	// ParseInLocation interprets the time as in the given location.
	// Second, when given a zone offset or abbreviation, Parse tries to match it
	// against the Local location; ParseInLocation uses the given location.
	t, err := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	if err != nil {
		log.Fatalln("ParseInLocation error:", err)
	}
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	t, err = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	if err != nil {
		log.Fatalln("ParseInLocation error:", err)
	}
	fmt.Println(t)
}

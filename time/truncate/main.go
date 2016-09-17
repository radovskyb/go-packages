package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t, err := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	if err != nil {
		log.Fatalln(err)
	}
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range trunc {
		// Truncate returns the result of rounding t down to a multiple
		// of d (since the zero time).
		// If d <= 0, Truncate returns t unchanged.
		fmt.Printf("t.Truncate(%6s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ParseTime parses a time header (such as the Date: header),
	// trying each of the three formats allowed by HTTP/1.1:
	// TimeFormat, time.RFC850, and time.ANSIC.
	t, err := http.ParseTime("Sat, 13 Feb 2016 17:30:00 GMT")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Time: %v\n", t)
}

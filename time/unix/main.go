package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	unixTimeStamp := "1432572732" // generate with Unix/Linux command >date +%s

	unixIntValue, err := strconv.ParseInt(unixTimeStamp, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	// Unix returns the local Time corresponding to the given Unix time,
	// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
	// It is valid to pass nsec outside the range [0, 999999999].
	// Not all sec values have a corresponding time value. One such
	// value is 1<<63-1 (the largest int64 value).
	timeStamp := time.Unix(unixIntValue, 0)

	fmt.Println(timeStamp)
}

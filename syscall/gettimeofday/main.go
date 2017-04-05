package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Create a new Timeval object.
	tv := &syscall.Timeval{}

	// Gettimeofday retrieves the time.
	if err := syscall.Gettimeofday(tv); err != nil {
		log.Fatalln(err)
	}

	// Print out the values from the Timeval type.
	fmt.Println(tv.Sec)
	fmt.Println(tv.Usec)
	fmt.Println(tv.Unix())
	fmt.Println(tv.Nano())
}

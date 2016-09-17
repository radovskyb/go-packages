package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Create a new file
	_, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Get the files last modified and access time stamps
	info, err := os.Stat("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the files last modified time
	fmt.Println(info.ModTime())

	// Wait one second before running the chtimes function
	time.Sleep(time.Second)

	// Chtimes changes the access and modification times of the named file,
	// similar to the Unix utime() or utimes() functions.
	//
	// The underlying filesystem may truncate or round the values to a less
	// precise time unit. If there is an error, it will be of type *PathError.
	if err := os.Chtimes("file.txt", time.Now(), time.Now()); err != nil {
		log.Fatalln(err)
	}

	// Print out the last modified time after changing it above
	info, err = os.Stat("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the time again after it's been updated
	fmt.Println(info.ModTime())
}

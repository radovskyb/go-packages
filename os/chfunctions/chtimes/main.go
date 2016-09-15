package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Create a new file
	os.Create("file.txt")

	// Get the files last modified and access time stamps
	info, _ := os.Stat("file.txt")

	// Print out the files last modified time
	fmt.Println(info.ModTime())

	// Wait one second before running the chtimes function
	time.Sleep(time.Second)

	// Chtimes changes the access and modification times of the named file,
	// similar to the Unix utime() or utimes() functions.
	//
	// The underlying filesystem may truncate or round the values to a less
	// precise time unit. If there is an error, it will be of type *PathError.
	os.Chtimes("file.txt", time.Now(), time.Now())

	// Print out the last modified time after changing it above
	info, _ = os.Stat("file.txt")

	// Print out the time again after it's been updated
	fmt.Println(info.ModTime())
}

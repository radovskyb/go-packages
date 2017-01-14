package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	str := "Hello, World!"

	// BytePtrFromString returns a pointer to a NUL-terminated array of
	// bytes containing the text of s. If s contains a NUL byte at any
	// location, it returns (nil, EINVAL).
	strBytesPtr, err := syscall.BytePtrFromString(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(strBytesPtr)
}

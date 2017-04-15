package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Open this file as read only and return it's file descriptor.
	//
	// Given a pathname for a file, Open returns a file descriptor, a
	// small, nonnegative integer for use in other system calls.
	// The file descriptor returned by a successful call will be the
	// lowest-numbered file descriptor not currently open for the process.
	//
	// See more at: http://man7.org/linux/man-pages/man2/open.2.html
	fd, err := syscall.Open("main.go", syscall.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file descriptor number.
	fmt.Println(fd)

	// Close the file descriptor.
	//
	// Close closes a file descriptor, so that it no longer refers to any
	// file and may be reused.
	//
	// See more at: http://man7.org/linux/man-pages/man2/close.2.html
	err = syscall.Close(fd)
	if err != nil {
		log.Fatalln(err)
	}
}

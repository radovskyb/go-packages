package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Getuid returns the real user ID of the calling process.
	fmt.Println(syscall.Getuid())
}

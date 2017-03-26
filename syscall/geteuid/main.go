package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Geteuid returns the effective user ID of the calling process.
	fmt.Println(syscall.Geteuid())
}

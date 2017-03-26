package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Getegid returns the effective group ID of the calling process.
	fmt.Println(syscall.Getegid())
}

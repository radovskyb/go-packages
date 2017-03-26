package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Getgid returns the real group ID of the calling process.
	fmt.Println(syscall.Getgid())
}

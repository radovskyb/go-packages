package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the host name reported by the kernel
	//
	// Hostname returns the host name reported by the kernel.
	hn, _ := os.Hostname()

	// Print the hostname
	fmt.Println(hn) // Benjamins-MacBook-Pro.local
}

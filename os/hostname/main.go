package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the host name reported by the kernel
	//
	// Hostname returns the host name reported by the kernel.
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the hostname
	fmt.Println(hn) // Benjamins-MacBook-Pro.local
}

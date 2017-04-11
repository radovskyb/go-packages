package main

import (
	"log"
	"syscall"
)

func main() {
	// Mkdir creates a new directory with the specified name and
	// permission bits.
	err := syscall.Mkdir("tmp", 0755)
	if err != nil {
		log.Fatalln(err)
	}
}

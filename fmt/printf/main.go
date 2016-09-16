package main

import (
	"fmt"
	"log"
)

func main() {
	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.
	n, err := fmt.Printf("Number: %d\n", 25)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Printed %d bytes above\n", n)
}

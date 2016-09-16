package main

import (
	"fmt"
	"log"
)

func main() {
	var first, second string

	// Sscanln is similar to Sscan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	n, err := fmt.Sscanln("First Second", &first, &second)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Scanned %d words\nFirst: %s\nSecond %s\n", n, first, second)
}

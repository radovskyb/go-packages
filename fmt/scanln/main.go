package main

import (
	"fmt"
	"log"
)

func main() {
	var Hello, World string

	fmt.Print("Please enter the text `Hello World!`: ")

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	_, err := fmt.Scanln(&Hello, &World)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Hello: %s\nWorld: %s\n", Hello, World)
}

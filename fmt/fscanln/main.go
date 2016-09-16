package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var Hello, World string

	// Fscanln is similar to Fscan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	n, err := fmt.Fscanln(strings.NewReader("Hello World!\nBlank"), &Hello, &World)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Scanned %d words\nHello: %s\nWorld: %s\n", n, Hello, World)
}

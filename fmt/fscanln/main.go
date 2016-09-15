package main

import (
	"fmt"
	"strings"
)

func main() {
	var Hello, World, Blank string

	// Fscanln is similar to Fscan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Fscanln(strings.NewReader("Hello World!\nBlank"), &Hello, &World, &Blank)

	fmt.Printf("Hello: %s\nWorld: %s\n", Hello, World)
	if Blank == "" {
		fmt.Println("Fscanln didn't scan into &Blank since a \\n " +
			"character came before it")
	}
}

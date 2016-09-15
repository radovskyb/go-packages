package main

import "fmt"

func main() {
	var Hello, World, Blank string

	fmt.Print("Please enter the text `Hello World!`: ")

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln(&Hello, &World, &Blank)

	fmt.Printf("Hello: %s\nWorld: %s\n", Hello, World)
	if Blank == "" {
		fmt.Println("Fscanln didn't scan into &Blank since a newline occured")
	}
}

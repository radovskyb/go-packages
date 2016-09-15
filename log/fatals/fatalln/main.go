package main

import "log"

func main() {
	// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
	log.Fatalln("Fatalln")
}

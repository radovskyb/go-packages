package main

import "log"

func main() {
	// Printf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	log.Printf("Print%s", "f")
}

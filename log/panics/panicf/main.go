package main

import "log"

func main() {
	// Panicf is equivalent to Printf() followed by a call to panic().
	log.Panicf("Panic%s", "f")
}

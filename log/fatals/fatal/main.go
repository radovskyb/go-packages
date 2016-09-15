package main

import "log"

func main() {
	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	log.Fatal("Fatal")
}

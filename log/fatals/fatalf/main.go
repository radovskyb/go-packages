package main

import "log"

func main() {
	// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
	log.Fatalf("Fatal%s", "f")
}

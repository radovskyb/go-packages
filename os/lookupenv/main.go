package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Lookup PATH in environment variables by its name using os.LookupEnv,
	// which also returns a bool if the key exists or not.
	//
	// LookupEnv retrieves the value of the environment variable named
	// by the key. If the variable is present in the environment the
	// value (which may be empty) is returned and the boolean is true.
	// Otherwise the returned value will be empty and the boolean will be false.
	PATH, found := os.LookupEnv("PATH")

	// Check if PATH exists as an environment variable and if so, print
	// it's returned string PATH, else print an error message
	if found {
		fmt.Println(PATH)
	} else {
		log.Fatalln("No PATH environment variable found.")
	}
}

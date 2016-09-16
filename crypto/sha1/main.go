package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

func main() {
	// Create a new sha1 hash.Hash object `h`
	//
	// New returns a new hash.Hash computing the SHA1 checksum.
	h := sha1.New()

	// Write a string to `h`
	_, err := h.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha1 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h.Sum(nil))

	// Hash and print out a sha1 hash from the same string
	// as above all in 1 line
	fmt.Printf("%x\n", sha1.Sum([]byte("Password123")))
}

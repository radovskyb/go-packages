package main

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	// SHA256:

	fmt.Println("SHA256")

	// Create a new sha256 hash.Hash object `h1`
	//
	// New returns a new hash.Hash computing the SHA256 checksum.
	h1 := sha256.New()

	// Write a string to `h1`
	_, err := h1.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha256 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h1.Sum(nil))

	// Hash and print out a sha256 hash from the same string
	// as above all in 1 line
	fmt.Printf("%x\n", sha256.Sum256([]byte("Password123")))

	// SHA224:

	fmt.Println("\nSHA224")

	// Create a new sha224 hash.Hash object `h2`
	//
	// New224 returns a new hash.Hash computing the SHA224 checksum.
	h2 := sha256.New224()

	// Write a string to `h2`
	_, err = h2.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha224 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h2.Sum(nil))

	// Hash and print out a sha224 hash from the same string
	// as above all in 1 line
	fmt.Printf("%x\n", sha256.Sum224([]byte("Password123")))
}

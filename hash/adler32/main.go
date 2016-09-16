package main

import (
	"fmt"
	"hash/adler32"
	"log"
)

func main() {
	// Create a new adler32 hashing object `hasher`
	//
	// New returns a new hash.Hash32 computing the Adler-32 checksum.
	// Its Sum method will lay the value out in big-endian byte order.
	hasher := adler32.New()

	// Write a []byte to `hasher`
	_, err := hasher.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the encrypted []byte from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", hasher.Sum(nil))

	// Complete the above hashing task above in one line below
	// by calling adler32.Checksum with the above []byte
	//
	// Checksum returns the Adler-32 checksum of data.
	fmt.Printf("%x\n", adler32.Checksum([]byte("Password123")))
}

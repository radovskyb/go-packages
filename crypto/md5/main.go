package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func main() {
	// Create a new md5 hash.Hash object `h`
	//
	// New returns a new hash.Hash computing the MD5 checksum.
	h := md5.New()

	// Write a string into `h`
	_, err := io.WriteString(h, "Password123")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the md5 encrypted string from above
	fmt.Printf("%x\n", h.Sum(nil))

	// Create a password string to hash using md5 hasing
	password := "Password123"

	// Hash the password in 1 step by calling md5.Sum and passing it
	// the password string as opposed to the above, first creating a
	// new md5 hasher, writing to it, and then calling md5.Sum with
	// nil as it's parameter
	fmt.Printf("%x", md5.Sum([]byte(password)))
}

package main

import (
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	// SHA512:

	fmt.Println("SHA512")

	// Create a new sha512 hash.Hash object `h1`
	//
	// New returns a new hash.Hash computing the SHA-512 checksum.
	h1 := sha512.New()

	// Write a string to `h1`
	_, err := h1.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha512 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h1.Sum(nil))

	// Hash and print out a sha512 hash from the same string
	// as above all in 1 line
	fmt.Printf("\n%x\n", sha512.Sum512([]byte("Password123")))

	// SHA512_224:

	fmt.Println("\nSHA512_224")

	// Create a new sha512_224 hash.Hash object `h2`
	//
	// New512_224 returns a new hash.Hash computing the SHA-512/224 checksum.
	h2 := sha512.New512_224()

	// Write a string to `h2`
	_, err = h2.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha512_224 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h2.Sum(nil))

	// Hash and print out a sha512_224 hash from the same string
	// as above all in 1 line
	fmt.Printf("%x\n", sha512.Sum512_224([]byte("Password123")))

	// SHA512_256:

	fmt.Println("\nSHA512_256")

	// Create a new sha512_256 hash.Hash object `h3`
	//
	// New512_256 returns a new hash.Hash computing the SHA-512/256 checksum.
	h3 := sha512.New512_256()

	// Write a string to `h3`
	_, err = h3.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha512_256 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h3.Sum(nil))

	// Hash and print out a sha512_256 hash from the same string
	// as above all in 1 line
	fmt.Printf("%x\n", sha512.Sum512_256([]byte("Password123")))

	// SHA384:

	fmt.Println("\nSHA384")

	// Create a new sha384 hash.Hash object `h4`
	//
	// New384 returns a new hash.Hash computing the SHA-384 checksum.
	h4 := sha512.New384()

	// Write a string to `h4`
	_, err = h4.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the sha384 encrypted string from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", h4.Sum(nil))

	// Hash and print out a sha384 hash from the same string
	// as above all in 1 line
	fmt.Printf("\n%x\n", sha512.Sum384([]byte("Password123")))
}

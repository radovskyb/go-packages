package main

import (
	"fmt"
	"hash/fnv"
	"log"
)

func main() {
	// fnv.New32:

	fmt.Println("fnv.New32")

	// Create a new fnv32 hashing object `hasher1`
	//
	// New32 returns a new 32-bit FNV-1 hash.Hash.
	hasher1 := fnv.New32()

	// Write a []byte to `hasher1`
	_, err := hasher1.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the encrypted []byte from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", hasher1.Sum32())

	// fnv.New32a:

	fmt.Println("\nfnv.New32a")

	// Create a new fnv32a hashing object `hasher2`
	//
	// New32a returns a new 32-bit FNV-1a hash.Hash.
	hasher2 := fnv.New32a()

	// Write a []byte to `hasher2`
	_, err = hasher2.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the encrypted []byte from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", hasher2.Sum32())

	// fnv.New64:

	fmt.Println("\nfnv.New64")

	// Create a new fnv64 hashing object `hasher3`
	//
	// New64 returns a new 64-bit FNV-1 hash.Hash.
	hasher3 := fnv.New64()

	// Write a []byte to `hasher3`
	_, err = hasher3.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the encrypted []byte from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", hasher3.Sum64())

	// fnv.New64a:

	fmt.Println("\nfnv.New64a")

	// Create a new fnv64a hashing object `hasher4`
	//
	// New64a returns a new 64-bit FNV-1a hash.Hash.
	hasher4 := fnv.New64a()

	// Write a []byte to `hasher4`
	_, err = hasher4.Write([]byte("Password123"))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the encrypted []byte from above
	//
	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	fmt.Printf("%x\n", hasher4.Sum64())
}

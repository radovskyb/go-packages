package main

import (
	"fmt"
	"hash/crc64"
	"log"
)

// http://golang.org/pkg/hash/crc64/#pkg-constants
//
// MakeTable returns the Table constructed from the specified polynomial.
//
// The ECMA polynomial, defined in ECMA 182.
var crcTable = crc64.MakeTable(crc64.ECMA)

func main() {
	// Checksum:

	// Checksum returns the CRC-64 checksum of data
	// using the polynomial represented by the Table.
	checksum64 := crc64.Checksum([]byte("Hello, World!"), crcTable)

	// Print out the checksum
	fmt.Printf("Checksum: %x\n", checksum64)

	// New:

	hasher := crc64.New(crcTable)

	_, err := hasher.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nHash64 value: %x\n", hasher.Sum64())

	// Reset resets the Hash to its initial state.
	hasher.Reset()

	_, err = hasher.Write([]byte("Hello!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Hash64 value: %x\n", hasher.Sum64())

	hasher.Reset()

	_, err = hasher.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Hash64 value: %x\n", hasher.Sum64())

	// Update:

	checksum := crc64.Checksum([]byte("Hello"), crcTable)

	// Print out `checksum`
	fmt.Printf("\nChecksum: %x\n", checksum)

	// Update the checksum
	updatedcsum := crc64.Update(checksum, crcTable, []byte(", World!"))

	// Print out the updated checksum
	fmt.Printf("Updated checksum: %x\n", updatedcsum)
}

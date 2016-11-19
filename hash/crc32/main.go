package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
)

// see http://golang.org/pkg/hash/crc32/#pkg-constants
//
// MakeTable returns the Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
//
// Castagnoli's polynomial, used in iSCSI.
// Has better error detection characteristics than IEEE.
// http://dx.doi.org/10.1109/26.231911
var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

func main() {
	// New:
	//
	// New creates a new hash.Hash32 computing the CRC-32 checksum
	// using the polynomial represented by the Table.
	crc := crc32.New(castagnoliTable)

	_, err := crc.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sum32: %x\n", crc.Sum32())

	// Checksum:
	//
	// Print out the same as the above with a single
	// line by calling crc32.Checksum
	fmt.Printf("Checksum: %x\n", crc32.Checksum([]byte("Hello, World!"), castagnoliTable))

	// NewIEEE:
	crcIEEE := crc32.NewIEEE()
	_, err = crcIEEE.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("crcIEEE Sum32: %x\n", crcIEEE.Sum32())

	// Reset resets the Hash to its initial state.
	crcIEEE.Reset()

	_, err = io.WriteString(crcIEEE, "Hello") // alternative way to write
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Sum32 after reset and new write: %x\n", crcIEEE.Sum32())

	// ChecksumIEEE:
	//
	// Print the same as the above's first write with a
	// single line by calling crc32.ChecksumIEEE
	//
	// ChecksumIEEE returns the CRC-32 checksum of data
	// using the IEEE polynomial.
	checksum := crc32.ChecksumIEEE([]byte("Hello, World!"))
	fmt.Printf("ChecksumIEEE: %x \n", checksum)

	// Update:
	//
	// Original checksum
	checksum = crc32.Checksum([]byte("Hello"), castagnoliTable)

	// Print out the original checksum
	fmt.Printf("\nChecksum: %x\n", checksum)

	// Update the checksum
	//
	// Update returns the result of adding the bytes in p to the crc.
	updatedchecksum := crc32.Update(checksum, castagnoliTable, []byte(", World!"))

	// Print out the updated checksum
	fmt.Printf("Updated Checksum : %x \n", updatedchecksum)
}

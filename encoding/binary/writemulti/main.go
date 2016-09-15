package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes buffer buf
	buf := new(bytes.Buffer)

	// Create a new slice of type interface{}, `data`,
	// then fill it with 3 different types of integer data
	// one uint16 type, one int8 type and one uint8 type.
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(245),
	}

	// Iterate through each of the data items
	for _, v := range data {
		// Print out the size of the bytes that binary.Write would
		// generate to encode each of the data items in the data slice
		fmt.Printf("%v ", binary.Size(v))

		// Encode and write v in the format of
		// binary.LittleEndian, to bytes buffer buf
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			log.Fatalf("binary.Write failed to encode %v: %v\n", v, err)
		}
	}

	// Print out the encoded bytes stored in bytes buffer buf
	fmt.Printf("\n% x\n", buf.Bytes())
}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)

	// Read reads structured binary data from r into data.
	// Data must be a pointer to a fixed-size value or a slice
	// of fixed-size values.
	// Bytes read from r are decoded using the specified byte order
	// and written to successive fields of the data.
	// When reading into structs, the field data for fields with
	// blank (_) field names is skipped; i.e., blank field names
	// may be used for padding.
	// When reading into a struct, all non-blank fields must be exported.
	//
	// LittleEndian is the little-endian implementation of ByteOrder.
	if err := binary.Read(buf, binary.LittleEndian, &pi); err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)
}

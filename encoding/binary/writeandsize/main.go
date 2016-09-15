package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi

	// Write writes the binary representation of data into w.
	// Data must be a fixed-size value or a slice of fixed-size
	// values, or a pointer to such data.
	// Bytes written to w are encoded using the specified byte order
	// and read from successive fields of the data.
	// When writing structs, zero values are written for fields
	// with blank (_) field names.
	if err := binary.Write(buf, binary.LittleEndian, pi); err != nil {
		log.Fatalln("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buf.Bytes())

	// Print out the size of the bytes that binary.Write would generate
	// to encode the float 123.456
	//
	// Size returns how many bytes Write would generate to encode
	// the value v, which must be a fixed-size value or a slice of
	// fixed-size values or a pointer to such data.
	// If v is neither of these, Size returns -1.
	fmt.Println(binary.Size(123.456)) // 8
}

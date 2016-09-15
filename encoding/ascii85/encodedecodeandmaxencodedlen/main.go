package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	// Create a new []byte
	byteSlice := []byte("I am a string! Hello, World!")

	// Create a new []byte buffer with the len of
	// ascii85.MaxEncodedLen for the len of byteSlice
	//
	// MaxEncodedLen returns the maximum length of an
	// encoding of n source bytes.
	buffer := make([]byte, ascii85.MaxEncodedLen(len(byteSlice)))

	// Encode the byteSlice []byte into the buffer []byte using
	// ascii85.Encode and return the number of bytes written
	//
	// Encode encodes src into at most MaxEncodedLen(len(src))
	// bytes of dst, returning the actual number of bytes written.
	//
	// The encoding handles 4-byte chunks, using a special encoding
	// for the last fragment, so Encode is not appropriate for use on
	// individual blocks of a large data stream.  Use NewEncoder() instead.
	//
	// Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
	// Encode does not add these.
	numEncodedBytes := ascii85.Encode(buffer, byteSlice)

	fmt.Println("Encoded", numEncodedBytes, "bytes")

	fmt.Printf("Encoded byteSlice: %s\n", buffer)

	newBuffer := make([]byte, len(buffer))

	// Now decode buffer which is the encoded byteSlice into a
	// newBuffer []byte and print the decoded string
	//
	// Decode decodes src into dst, returning both the number
	// of bytes written to dst and the number consumed from src.
	// If src contains invalid ascii85 data, Decode will return the
	// number of bytes successfully written and a CorruptInputError.
	// Decode ignores space and control characters in src.
	// Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
	// Decode expects these to have been stripped by the caller.
	//
	// If flush is true, Decode assumes that src represents the
	// end of the input stream and processes it completely rather
	// than wait for the completion of another 32-bit block.
	//
	// NewDecoder wraps an io.Reader interface around Decode.
	_, _, err := ascii85.Decode(newBuffer, buffer, true)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Print the string from newBuffer decoded from buffer
	fmt.Printf("%s\n", newBuffer)
}

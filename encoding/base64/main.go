package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	msg := "Hello, World!"

	// StdEncoding is the standard base64 encoding, as defined in
	// RFC 4648.
	//
	// EncodeToString returns the base64 encoding of src.
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	// DecodedLen returns the maximum length in bytes of the decoded data
	// corresponding to n bytes of base64-encoded data.
	fmt.Println(base64.StdEncoding.DecodedLen(len(encoded)))

	// DecodeString returns the bytes represented by the base64 string s.
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}

	fmt.Println(string(decoded))

	fmt.Println()

	var buf bytes.Buffer
	// NewEncoder returns a new base64 stream encoder.  Data written to
	// the returned writer will be encoded using enc and then written to w.
	// Base64 encodings operate in 4-byte blocks; when finished
	// writing, the caller must Close the returned encoder to flush any
	// partially written blocks.
	encoder := base64.NewEncoder(
		base64.StdEncoding,
		&buf,
	)
	_, err = encoder.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}
	if err := encoder.Close(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(buf.String())

	// NewDecoder constructs a new base64 stream decoder.
	decoder := base64.NewDecoder(
		base64.StdEncoding,
		&buf,
	)

	_, err = io.Copy(os.Stdout, decoder)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n")

	str := "Hello, World!"
	// EncodedLen returns the length in bytes of the base64 encoding
	// of an input buffer of length n.
	bytesBuffer := make([]byte, base64.StdEncoding.EncodedLen(len(str)))

	// Encode encodes src using the encoding enc, writing
	// EncodedLen(len(src)) bytes to dst.
	//
	// The encoding pads the output to a multiple of 4 bytes,
	// so Encode is not appropriate for use on individual blocks
	// of a large data stream.  Use NewEncoder() instead.
	base64.StdEncoding.Encode(bytesBuffer, []byte(str))
	fmt.Println(string(bytesBuffer))

	decodedBuffer := make([]byte, base64.StdEncoding.DecodedLen(len(bytesBuffer)))
	// Decode decodes src using the encoding enc.  It writes at most
	// DecodedLen(len(src)) bytes to dst and returns the number of bytes
	// written.  If src contains invalid base64 data, it will return the
	// number of bytes successfully written and CorruptInputError.
	// New line characters (\r and \n) are ignored.
	base64.StdEncoding.Decode(decodedBuffer, bytesBuffer)
	fmt.Println(string(decodedBuffer))
}

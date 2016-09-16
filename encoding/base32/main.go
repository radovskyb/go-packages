package main

import (
	"encoding/base32"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := "Hello, World!"
	// EncodedLen returns the length in bytes of the base32 encoding
	// of an input buffer of length n.
	fmt.Println(base32.StdEncoding.EncodedLen(len(str)))

	// EncodeToString returns the base32 encoding of src.
	encoded := base32.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encoded)

	// DecodeString returns the bytes represented by the base32 string s.
	decoded, err := base32.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatalln(err)
	}

	// DecodedLen returns the maximum length in bytes of the decoded data
	// corresponding to n bytes of base32-encoded data.
	fmt.Println(base32.StdEncoding.DecodedLen(len(encoded)))

	fmt.Println(string(decoded))
	fmt.Println()

	// NewEncoder returns a new base32 stream encoder.  Data written to
	// the returned writer will be encoded using enc and then written to w.
	// Base32 encodings operate in 5-byte blocks; when finished
	// writing, the caller must Close the returned encoder to flush any
	// partially written blocks.
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	_, err = encoder.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}
	if err := encoder.Close(); err != nil {
		log.Fatalln(err)
	}

	// NewReader returns a new Reader reading from s.
	// It is similar to bytes.NewBufferString but more efficient
	// and read-only.
	sr := strings.NewReader("JBSWY3DPFQQFO33SNRSCC===")

	// NewDecoder constructs a new base32 stream decoder.
	decoder := base32.NewDecoder(base32.StdEncoding, sr)

	fmt.Println()

	_, err = io.Copy(os.Stdout, decoder)
	if err != nil {
		log.Fatalln(err)
	}
}

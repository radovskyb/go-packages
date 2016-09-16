package main

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Create a new bytes buffer buf
	var buf bytes.Buffer

	// Create a new ascii85 encoder and pass it &buf to encode into
	//
	// NewEncoder returns a new ascii85 stream encoder. Data written to
	// the returned writer will be encoded and then written to w.
	// Ascii85 encodings operate in 32-bit blocks; when finished
	// writing, the caller must Close the returned encoder to flush any
	// trailing partial block.
	encoderstream := ascii85.NewEncoder(&buf)

	// Write a byte slice to the encoderstream
	_, err := encoderstream.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	// Close the stream now that writing is finished
	if err := encoderstream.Close(); err != nil {
		log.Fatalln(err)
	}

	// Print out the encoded string inside of buf
	fmt.Println(buf.String())

	// Create a new string
	var str string

	// Store the encoded string in buf into str
	str = buf.String()

	// Create a new ascii85 decoder and pass it a new strings.Reader
	// to read from str
	decoderstream := ascii85.NewDecoder(strings.NewReader(str))

	// Read all of the contents from the decoderstream into decodedstr
	decodedstr, err := ioutil.ReadAll(decoderstream)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the decoded string from decoderstream
	// stored in decodedstr
	fmt.Printf("%s\n", decodedstr)
}

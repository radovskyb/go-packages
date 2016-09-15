package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
)

func main() {
	// A WordDecoder decodes MIME headers containing RFC 2047 encoded-words.
	dec := new(mime.WordDecoder)

	// Decode decodes an RFC 2047 encoded-word.
	header, err := dec.Decode("=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(header) // ¡Hola, señor!

	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// charset into UTF-8.
	// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
	// are handled by default.
	// One of the the CharsetReader's result values must be non-nil.
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("Unhandled charset %q", charset)
		}
	}

	header, err = dec.Decode("=?x-case?q?hello!?=")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(header)
}

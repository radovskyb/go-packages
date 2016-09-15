package main

import (
	"fmt"
	"mime"
)

func main() {
	// A WordEncoder is a RFC 2047 encoded-word encoder.
	enc := new(mime.WordEncoder)

	// Encode returns the encoded-word form of s. If s is ASCII without special
	// characters, it is returned unchanged. The provided charset is the IANA
	// charset name of s. It is case insensitive.
	fmt.Println(enc.Encode("utf-8", "¡Hola, señor!"))

	// QEncoding represents the Q-encoding scheme as defined by RFC 2047.
	fmt.Println(mime.QEncoding.Encode("utf-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("utf-8", "Hello!"))
	fmt.Println(mime.QEncoding.Encode("ISO-8859-1", "Caf\xE9"))

	// BEncoding represents Base64 encoding scheme as defined by RFC 2045.
	fmt.Println(mime.BEncoding.Encode("UTF-8", "¡Hola, señor!"))
}

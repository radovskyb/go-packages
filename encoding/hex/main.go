package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new byte slice b
	b := []byte("Hello, World!")

	// Create a new byte slice encodeBuf, which is the size of
	// the encoded length of the hex encoding of the length of b
	//
	// EncodedLen returns the length of an encoding of n source bytes.
	encodeBuf := make([]byte, hex.EncodedLen(len(b)))

	// Encode b and store it's encoded result into encodeBuf
	//
	// Encode encodes src into EncodedLen(len(src))
	// bytes of dst.  As a convenience, it returns the number
	// of bytes written to dst, but this value is always
	// EncodedLen(len(src)).
	// Encode implements hexadecimal encoding.
	hex.Encode(encodeBuf, b)

	// Print out buf as a string
	fmt.Println(string(encodeBuf))

	// Create a new string, strBufEncode and store the result of
	// calling hex.EncodeToString with byte slice b
	//
	// EncodeToString returns the hexadecimal encoding of src.
	strBufEncode := hex.EncodeToString(b)

	// Print out strBufEncode
	fmt.Println(strBufEncode)

	// Create a new byte slice decodeBuf to store the decoded
	// version of strBufEncode which is currently a hex string. It's size
	// is the length that decoding strBufEncode will be
	//
	// DecodedLen returns the length of a decoding of n hex source bytes.
	decodeBuf := make([]byte, hex.DecodedLen(len(strBufEncode)))

	// Decode strBufEncode as a byte slice into decodeBuf
	//
	// Decode decodes src into DecodedLen(len(src)) bytes,
	// returning the actual number of bytes written to dst.
	//
	// If Decode encounters invalid input, it returns an
	// error describing the failure.
	decoded, err := hex.Decode(decodeBuf, []byte(strBufEncode))
	if err != nil {
		log.Fatalln(err)
	}

	// Print out decodeBuf as a string
	fmt.Printf("Decoded %d bytes: %s\n", decoded, decodeBuf)

	// Now decode strBufEncode into a []byte using hex.DecodeString
	//
	// DecodeString returns the bytes represented by the
	//hexadecimal string s.
	bufDecode, err := hex.DecodeString(strBufEncode)
	if err != nil {
		fmt.Println("Could not decode strBufEncode:", err)
	}

	// Print bufDecode as a string
	fmt.Println(string(bufDecode))

	// Create a hex dump from byte slice b
	//
	// Dump returns a string that contains a hex dump of the
	// given data. The format of the hex dump matches the
	// output of `hexdump -C` on the command line.
	dumpedStr := hex.Dump(b)

	// Print dumpedStr
	fmt.Println(dumpedStr)

	// Create a hex.Dumper that will write hex dumps to os.Stdout
	// when bytes are written to it
	//
	// Dumper returns a WriteCloser that writes a hex dump
	// of all written data to w. The format of the dump matches
	// the output of `hexdump -C` on the command line.
	dumper := hex.Dumper(os.Stdout)

	// Write a byte slice to hex.Dumper dumper which
	// will print a hex dump of the byte slice to os.Stdout
	_, err = dumper.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dumper now that we are finished with it
	if err := dumper.Close(); err != nil {
		log.Fatalln(err)
	}
}

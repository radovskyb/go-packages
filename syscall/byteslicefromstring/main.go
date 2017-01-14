package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	str := "Hello, World!"

	// ByteSliceFromString returns a NUL-terminated slice of bytes
	// containing the text of s. If s contains a NUL byte at any
	// location, it returns (nil, EINVAL).
	strBytes, err := syscall.ByteSliceFromString(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(strBytes)
	fmt.Println(len(strBytes) == len(str)) // false
	fmt.Println(strBytes[len(strBytes)-1]) // NUL-terminator
}

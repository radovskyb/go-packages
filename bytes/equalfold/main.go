package main

import (
	"bytes"
	"fmt"
)

func main() {
	// EqualFold reports whether s and t, interpreted as
	// UTF-8 strings, are equal under Unicode case-folding.
	if bytes.EqualFold([]byte{'a'}, []byte{97}) {
		fmt.Println("[byte slice with and byte slice with 97]: They are the same.")
	}
}

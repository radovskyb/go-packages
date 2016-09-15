package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Equal returns a boolean reporting whether a and b are
	// the same length and contain the same bytes. A nil argument
	// is equivalent to an empty slice.
	if bytes.Equal([]byte{'a'}, []byte{'a'}) {
		fmt.Println("[a and a byte slices]: They are the same.")
	}

	if !bytes.Equal([]byte{'a'}, []byte{'b'}) {
		fmt.Println("[a and b byte slices]: They are not the same.")
	}

	// Here this one also returns true since as it states above,
	// a nil argument is equivalent to an empty slice.
	if bytes.Equal([]byte{}, nil) {
		fmt.Println("[empty byte slice and nil]: They are the same.")
	}
}

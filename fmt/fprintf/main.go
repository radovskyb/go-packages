package main

import (
	"fmt"
	"os"
)

type Object struct {
	Field int
}

func main() {
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(os.Stdout, "Object: %#v\n", &Object{5})
}

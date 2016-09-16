package main

import (
	"fmt"
	"log"
	"os"
)

type Object struct {
	Field int
}

func main() {
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	_, err := fmt.Fprintf(os.Stdout, "Object: %#v\n", &Object{5})
	if err != nil {
		log.Fatalln(err)
	}
}

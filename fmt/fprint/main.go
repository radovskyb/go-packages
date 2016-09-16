package main

import (
	"fmt"
	"log"
	"os"
)

type Object struct {
	Field int
}

func (o *Object) String() string {
	return fmt.Sprintf("Object: %d\n", o.Field)
}

func main() {
	// Fprint formats using the default formats for its operands and writes to w.
	// Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error encountered.
	_, err := fmt.Fprint(os.Stdout, &Object{5})
	if err != nil {
		log.Fatalln(err)
	}
}

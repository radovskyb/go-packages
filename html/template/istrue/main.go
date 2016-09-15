package main

import (
	"fmt"
	"html/template"
)

func main() {
	var b bool

	// IsTrue reports whether the value is 'true', in the sense of
	// not the zero of its type, and whether the value has a
	// meaningful truth value. This is the definition of
	// truth used by if and other such actions.
	truth, okay := template.IsTrue(b)
	fmt.Println(truth, okay)

	b = true
	truth, okay = template.IsTrue(b)
	fmt.Println(truth, okay)
}

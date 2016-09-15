package main

import (
	"fmt"
	"strings"
)

type Object struct {
	FieldOne int
	FieldTwo string
}

func main() {
	obj := new(Object)

	var i int

	// Fscan scans text read from r, storing successive space-separated
	// values into successive arguments.  Newlines count as space.  It
	// returns the number of items successfully scanned.  If that is less
	// than the number of arguments, err will report why.
	fmt.Fscan(strings.NewReader("5 String 1"), &obj.FieldOne, &obj.FieldTwo, &i)

	fmt.Println(obj.FieldOne)
	fmt.Println(obj.FieldTwo)
	fmt.Println(i)
}

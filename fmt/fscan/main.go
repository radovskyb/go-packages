package main

import (
	"fmt"
	"log"
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
	_, err := fmt.Fscan(strings.NewReader("5 String 1"), &obj.FieldOne, &obj.FieldTwo, &i)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(obj.FieldOne)
	fmt.Println(obj.FieldTwo)
	fmt.Println(i)
}

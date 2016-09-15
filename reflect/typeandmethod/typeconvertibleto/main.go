package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := "I'm a string"

	t1 := reflect.TypeOf(a) // int
	t2 := reflect.TypeOf(b) // string

	// Print whether an int can be converted to a string
	fmt.Println(t1.ConvertibleTo(t2)) // true

	// Print whether a string can be converted to an int
	fmt.Println(t2.ConvertibleTo(t1)) // false
}

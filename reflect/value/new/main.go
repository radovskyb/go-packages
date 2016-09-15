package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a new integer pointer using reflect
	//
	// New returns a Value representing a pointer to a new zero value
	// for the specified type. That is, the returned Value's Type
	// is PtrTo(typ).
	intPtr := reflect.New(reflect.TypeOf((*int)(nil)).Elem())

	// Set intPtr's underlying value
	intPtr.Elem().SetInt(25)

	// Print intPtr's underlying value
	fmt.Println(intPtr.Elem().Int())
}

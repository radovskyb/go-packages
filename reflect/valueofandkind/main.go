package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10

	// Return a new reflect.Value type initialized to the concrete
	// value stored in a
	//
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i.  ValueOf(nil) returns the zero Value.
	aValue := reflect.ValueOf(a)

	// Get a's kind from `aValue`
	//
	// A Kind represents the specific kind of type that a Type represents.
	// The zero Kind is not a valid kind.
	//
	// Kind returns v's Kind.
	// If v is the zero Value (IsValid returns false), Kind returns
	// Invalid.
	aKind := aValue.Kind()

	// Print out aKind (The kind of variable a); aKind.String()
	fmt.Println(aKind) // int
}

package main

import (
	"fmt"
	"reflect"
)

type myStringerInterface interface {
	String() string
}

type myStringerType int

func (m myStringerType) String() string {
	return string(m)
}

func main() {
	// Create a new myStringerType object `ms`
	ms := myStringerType(10)

	// Create a new reflect.Type object `typeofms` from `ms`
	typeofms := reflect.TypeOf(ms)

	// Print out whether or not `ms` implements
	// the myStringerInterface interface
	fmt.Println(typeofms.Implements(
		// Elem returns a type's element type.
		reflect.TypeOf((*myStringerInterface)(nil)).Elem()),
	) // true
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a string slice str
	var str []string

	// Cteate v variable of reflect.Value object type from &str
	var v reflect.Value = reflect.ValueOf(&str)

	// Make strval to adopt to element's type - in this
	// case a string type
	v = v.Elem()

	// Append some strings as reflect.Value type objects to v
	//
	// Append appends the values x to a slice s and returns the
	// resulting slice. As in Go, each x's value must be assignable
	// to the slice's element type.
	v = reflect.Append(v, reflect.ValueOf("abc"))
	v = reflect.Append(v, reflect.ValueOf("def"))
	v = reflect.Append(v, reflect.ValueOf("xyz"))

	// Print v.Kind() and v.Type()
	// Prints: `v is a type of slice and is a []string`
	fmt.Println("v is a type of:", v.Kind(), "and is a", v.Type())

	// Print v's elements
	//
	// Interface returns v's current value as an interface{}.
	// It is equivalent to:
	//	var i interface{} = (v's underlying value)
	// It panics if the Value was obtained by accessing
	// unexported struct fields.
	fmt.Println("v's elements:", v.Interface())

	// Append a reflect.Value of type []string with 2 string elements
	// to v
	//
	// AppendSlice appends a slice t to a slice s and returns
	// the resulting slice. The slices s and t must have the
	// same element type.
	v = reflect.AppendSlice(v, reflect.ValueOf([]string{"mno", "qrs"}))

	// Now print v's elements again after appending to it
	fmt.Println("v's elements after reflect.AppendSlice:", v.Interface())
}

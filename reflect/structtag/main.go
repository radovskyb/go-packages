package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	MyField string `json:"myfield" type:"myfieldtype" default:"0"`
}

func main() {
	// Create a new MyStruct struct type `ms`
	ms := MyStruct{}

	// Get `ms`'s type
	st := reflect.TypeOf(ms)

	// Get `st`'s first field (returns a reflect.StructField object)
	field := st.Field(0)

	// Get the field's tag (returns a reflect.StructTag object)
	//
	// A StructTag is the tag string in a struct field.
	//
	// By convention, tag strings are a concatenation of
	// optionally space-separated key:"value" pairs.
	// Each key is a non-empty string consisting of non-control
	// characters other than space (U+0020 ' '), quote (U+0022 '"'),
	// and colon (U+003A ':').  Each value is quoted using U+0022 '"'
	// characters and Go string literal syntax.
	tag := field.Tag

	// Print out the full tag
	fmt.Println(tag)

	// Now get & print the individual parts of the tag using tag.Get(key)
	//
	// Get returns the value associated with key in the tag string.
	// If there is no such key in the tag, Get returns the empty string.
	// If the tag does not have the conventional format, the value
	// returned by Get is unspecified.
	fmt.Println(tag.Get("json"))    // myfield
	fmt.Println(tag.Get("type"))    // myfieldtype
	fmt.Println(tag.Get("default")) // default
}

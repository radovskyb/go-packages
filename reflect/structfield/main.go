package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	MyField int
}

func main() {
	// Create a new MyStruct `myStruct`
	myStruct := &MyStruct{MyField: 5}

	// Create a new `myStructIndirect` from `myStruct`
	myStructIndirect := reflect.Indirect(reflect.ValueOf(myStruct))

	// Create a new reflect.StructField type `field` from
	// `myStructIndirect`'s first field.
	//
	// A StructField describes a single field in a struct.
	var field reflect.StructField = myStructIndirect.Type().Field(0)

	// Get the field's name
	fmt.Println(field.Name) // MyField

	// Get the field's type
	fmt.Println(field.Type) // int

	// Get the field's value
	fmt.Println(myStructIndirect.Field(0))
}

package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	FieldOne string
	Second
}

type Second struct {
	FieldOne int
}

func main() {
	// Create structVal, a reflect.Value object of type MyStruct{"Hello"}
	structVal := reflect.ValueOf(MyStruct{"Hello", Second{5}})

	fmt.Println("structVal.NumField:", structVal.NumField())

	fmt.Println()

	// Get the 0th field from the structVal struct
	//
	// Field returns the i'th field of the struct v.
	// It panics if v's Kind is not Struct or i is out of range.
	fmt.Println(structVal.Field(0))

	// Get the first field from the struct located at the second field
	// in the MyStruct structure
	//
	// FieldByIndex returns the nested field corresponding to index.
	// It panics if v's Kind is not struct.
	fmt.Println(structVal.FieldByIndex([]int{1, 0})) // 5

	// Get the 0th field by name
	//
	// FieldByName returns the struct field with the given name.
	// It returns the zero Value if no field was found.
	// It panics if v's Kind is not struct.
	fmt.Println(structVal.FieldByName("FieldOne"))

	// Create function f to define how to find a field in
	// structVal's struct
	f := func(s string) bool {
		if s == "FieldOne" {
			return true
		}
		return false
	}

	// Get a fieldVal from structVal by calling structVal.FieldByNameFunc
	// and passing it the f function created above
	//
	// FieldByNameFunc returns the struct field with a name
	// that satisfies the match function.
	// It panics if v's Kind is not struct.
	// It returns the zero Value if no field was found.
	var fieldVal = structVal.FieldByNameFunc(f)

	// Print out fieldVal
	fmt.Println(fieldVal)

	fmt.Println()

	var a interface{} = 10
	intVal := reflect.ValueOf(&a)

	// InterfaceData returns the interface v's value as a uintptr pair.
	// It panics if v's Kind is not Interface.
	fmt.Println(intVal.Elem().InterfaceData())

	fmt.Println()

	// IsNil reports whether its argument v is nil. The argument must be
	// a chan, func, interface, map, pointer, or slice value; if it is
	// not, IsNil panics. Note that IsNil is not always equivalent to a
	// regular comparison with nil in Go. For example, if v was created
	// by calling ValueOf with an uninitialized interface variable i,
	// i==nil will be true but v.IsNil will panic as v will be the zero
	// Value.
	fmt.Println("intVal.IsNil:", intVal.IsNil())

	fmt.Println()

	// IsValid reports whether v represents a value.
	// It returns false if v is the zero Value.
	// If IsValid returns false, all other methods except String panic.
	// Most functions and methods never return an invalid value.
	// If one does, its documentation states the conditions explicitly.
	fmt.Println("intVal.IsValid:", intVal.IsValid())
}

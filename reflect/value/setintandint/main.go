package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a new integer pointer using reflect.New
	intPtr := reflect.New(reflect.TypeOf((*int)(nil)).Elem())

	// Set intPtr's underlying value
	//
	// SetInt sets v's underlying value to x.
	// It panics if v's Kind is not Int, Int8, Int16, Int32,
	// or Int64, or if CanSet() is false.
	intPtr.Elem().SetInt(10)

	// Print intPtr's underlying value
	//
	// Int returns v's underlying value, as an int64.
	// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
	fmt.Println(intPtr.Elem().Int()) // 10

}

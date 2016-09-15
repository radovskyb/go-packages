package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a []int
	var b []int
	value1 := reflect.ValueOf(&a)
	value2 := reflect.ValueOf(&b)

	// Make a pointer to the memory address for a
	value1 = reflect.Indirect(value1)

	// Make a pointer to the memory address for b
	value2 = reflect.Indirect(value2)

	value2 = reflect.NewAt(value1.Type(), unsafe.Pointer(
		reflect.ValueOf(&b).Pointer(),
	)).Elem()

	b = append(b, 1)
	a = append(a, 2)

	fmt.Println(value1.Pointer(), value2.Pointer(), a, b)

	// Set value1's underlying data to value2's underlying data
	//
	// Set assigns x to the value v.
	// It panics if CanSet returns false.
	// As in Go, x's value must be assignable to v's type.
	value1.Set(value2)

	fmt.Println(
		value1.Pointer(),
		value2.Pointer(),
		a, b,
	)
}

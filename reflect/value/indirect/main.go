package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i []int                  // slice
	var str [3]string            // array
	var m = make(map[int]string) // map

	var iV = reflect.ValueOf(&i)
	fmt.Println("&i's type:", iV.Kind())

	// Indirect returns the value that v points to.
	// If v is a nil pointer, Indirect returns a zero Value.
	// If v is not a pointer, Indirect returns v.
	indirectI := reflect.Indirect(iV)

	// Here we are using the above reflect.Indirect to find out
	// what type i points to, since iV was a pointer and did not
	// print slice when asked iV.Kind()
	fmt.Println("i's type:", indirectI.Kind())

	fmt.Println()

	sV := reflect.ValueOf(&str)
	fmt.Println("&str's type:", sV.Kind())
	fmt.Println("str's type:", reflect.Indirect(sV).Kind())

	fmt.Println()

	mV := reflect.ValueOf(&m)
	fmt.Println("&m's type:", mV.Kind())
	fmt.Println("m's type:", reflect.Indirect(mV).Kind())
}

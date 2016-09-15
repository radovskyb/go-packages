package main

import (
	"fmt"
	"reflect"
)

type myInt struct {
	int
}

func main() {
	// Create a new myInt pointer
	m := &myInt{10}

	// Create a reflect.Type which is a pointer to a myInt object type
	myIntPtr := reflect.PtrTo(reflect.TypeOf(myInt{}))

	// Check and see if the type of myIntPtr and m's type are the same
	if myIntPtr == reflect.TypeOf(m) {
		fmt.Println("Their types are the same")
	}
}

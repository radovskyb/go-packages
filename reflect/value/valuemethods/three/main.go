package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int, 5)

	chanVal := reflect.ValueOf(ch)

	// Send sends x on the channel v.
	// It panics if v's kind is not Chan or if x's type is not
	// the same type as v's element type.
	// As in Go, x's value must be assignable to the
	//channel's element type.
	chanVal.Send(reflect.ValueOf(10))

	// Recv receives and returns a value from the channel v.
	// It panics if v's Kind is not Chan.
	// The receive blocks until a value is ready.
	// The boolean value ok is true if the value x corresponds to a send
	// on the channel, false if it is a zero value received because
	// the channel is closed.
	val, ok := chanVal.Recv()
	if ok {
		fmt.Println("Received", val, "from chanVal")
	}

	fmt.Println()

	// Create a new []int
	var intSlice []int

	// Create intVal, a reflect.Value object created from intSlice
	intVal := reflect.ValueOf(&intSlice)

	// Check if intVal can set and if not, assign it to
	// it's underlying element slice which can set
	//
	// CanSet reports whether the value of v can be changed.
	// A Value can be changed only if it is addressable and was not
	// obtained by the use of unexported struct fields.
	// If CanSet returns false, calling Set or any type-specific
	// setter (e.g., SetBool, SetInt) will panic.
	if !intVal.CanSet() {
		// Elem returns the value that the interface v contains
		// or that the pointer v points to.
		// It panics if v's Kind is not Interface or Ptr.
		// It returns the zero Value if v is nil.
		intVal = intVal.Elem()
	}

	// Set intVal's values to some []int values, 1 through to 5
	intVal.Set(reflect.ValueOf([]int{1, 2, 3, 4, 5}))

	// Print out intVal's current elements and length
	fmt.Println("intVal's elements before SetLen:", intVal.Interface())
	fmt.Println("intVal's length before SetLen:", intVal.Len())

	// Set the length of intVal to 3 from 5
	//
	// SetLen sets v's length to n.
	// It panics if v's Kind is not Slice or if n is negative or
	// greater than the capacity of the slice.
	intVal.SetLen(3)

	// Print out intVal's elements and length after calling SetLen
	fmt.Println("intVal's elements after SetLen:", intVal.Interface())
	fmt.Println("intVal's length after SetLen:", intVal.Len())
}

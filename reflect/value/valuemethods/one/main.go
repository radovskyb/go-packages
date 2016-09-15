package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	v := reflect.ValueOf(a)

	// Int returns v's underlying value, as an int64.
	// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
	fmt.Println(v.Int())

	fmt.Println()

	// CanAddr reports whether the value's address can be obtained
	// with Addr. Such values are called addressable. A value is
	// addressable if it is an element of a slice, an element of an
	// addressable array, a field of an addressable struct, or the
	// result of dereferencing a pointer.
	// If CanAddr returns false, calling Addr will panic.
	if v.CanAddr() {
		// Addr returns a pointer value representing the address of v.
		// It panics if CanAddr() returns false.
		// Addr is typically used to obtain a pointer to a struct field
		// or slice element in order to call a method that requires a
		// pointer receiver.
		fmt.Println(v.Addr())
	} else {
		fmt.Println("v can't Addr()")
		v := reflect.Indirect(reflect.ValueOf(&a))
		fmt.Println("Indirect of &a's Addr():", v.Addr())
	}

	fmt.Println()

	boolVal := reflect.ValueOf(true)

	// Bool returns v's underlying value.
	// It panics if v's kind is not Bool.
	fmt.Println(boolVal.Bool())

	fmt.Println()

	b := []byte{'a', 'b', 'c'}

	bytesVal := reflect.ValueOf(b)

	// Print out the byte reflect.Value object's bytes
	//
	// Bytes returns v's underlying value.
	// It panics if v's underlying value is not a slice of bytes.
	fmt.Println(string(bytesVal.Bytes()))

	fmt.Println()

	// Create a new function and store it in variable f. It takes in
	// a single int and returns a single int
	f1 := func(i int) int {
		fmt.Println("Calling Function, i:", i)
		return i
	}

	// Create funcValue, a reflect.Value of a function type
	funcValue1 := reflect.ValueOf(f1)

	// Call the function using funcValue.Call and pass it an
	// argument, a, the integer declared earlier, as a reflect.Value
	//
	// Call calls the function v with the input arguments in.
	// For example, if len(in) == 3, v.Call(in) represents the
	// Go call v(in[0], in[1], in[2]).
	// Call panics if v's Kind is not Func.
	// It returns the output results as Values.
	// As in Go, each input argument must be assignable to the
	// type of the function's corresponding input parameter.
	// If v is a variadic function, Call creates the variadic
	// slice parameter itself, copying in the corresponding values.
	result1 := funcValue1.Call([]reflect.Value{reflect.ValueOf(a)})
	fmt.Println(result1[0], "returned from funcValue1.Call")

	fmt.Println()

	// Create another function that this time takes an int slice
	// as a variadic args list of ints and returns an int slice
	f2 := func(intSlice ...int) []int {
		for _, v := range intSlice {
			fmt.Println(v)
		}
		return intSlice
	}

	// Create another function reflect.Value type object
	funcValue2 := reflect.ValueOf(f2)

	// Create an in variable of type []reflect.Value to pass to
	// funcValue2.CallSlice
	in := []reflect.Value{reflect.ValueOf([]int{1, 2, 3})}

	// Call f2 using funcValue2.Call and pass it the `in` variable
	//
	// CallSlice calls the variadic function v with the input arguments in,
	// assigning the slice in[len(in)-1] to v's final variadic argument.
	// For example, if len(in) == 3, v.CallSlice(in) represents the
	// Go call v(in[0], in[1], in[2]...).
	// CallSlice panics if v's Kind is not Func or if v is not variadic.
	// It returns the output results as Values.
	// As in Go, each input argument must be assignable to the
	// type of the function's corresponding input parameter.
	result2 := funcValue2.CallSlice(in)
	fmt.Println(result2[0], "returned from funcValue2.Call")

	fmt.Println()

	// CanInterface reports whether Interface can be used
	// without panicking.
	fmt.Println("bytesVal can interface:", bytesVal.CanInterface())

	fmt.Println()

	// CanSet reports whether the value of v can be changed.
	// A Value can be changed only if it is addressable and was not
	// obtained by the use of unexported struct fields.
	// If CanSet returns false, calling Set or any type-specific
	// setter (e.g., SetBool, SetInt) will panic.
	fmt.Println("bytesVal can set:", bytesVal.CanSet())

	v = reflect.Indirect(reflect.ValueOf(&a))
	fmt.Println("&a Inderect can set:", v.CanSet())

	fmt.Println()

	// Create sliceVal, a reflect.Value object of type []byte
	sliceVal := reflect.ValueOf(make([]byte, 10))

	// Print sliceVal's capacity
	//
	// Cap returns v's capacity.
	// It panics if v's Kind is not Array, Chan, or Slice.
	fmt.Println(sliceVal.Cap())

	fmt.Println()

	// Create chanVal, a reflect.Value object of type chan int
	chanVal := reflect.ValueOf(make(chan int, 1))

	// Close the channel chanVal
	//
	// Close closes the channel v.
	// It panics if v's Kind is not Chan.
	chanVal.Close()

	fmt.Println("chanVal was closed.")

	fmt.Println()

	// Create a new integer x set to 97
	x := 97

	// Create intVal1, a reflect.Value object of type int
	intVal1 := reflect.ValueOf(x)

	// Create stringVal, a reflect.Value object created by converting
	// intVal to type string
	//
	// Convert returns the value v converted to type t.
	// If the usual Go conversion rules do not allow conversion
	// of the value v to type t, Convert panics.
	stringVal := intVal1.Convert(reflect.TypeOf(""))

	// Print out stringVal now that it's a string
	fmt.Println(stringVal.Kind(), stringVal.String()) // sting a

	fmt.Println()

	intVal2 := reflect.ValueOf(&x)

	// Print out the value that intVal2 contains
	//
	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	fmt.Println(intVal2.Elem())
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Alignof takes an expression x of any type and returns the
	// required alignment of a hypothetical variable v as if v
	// was declared via var v = x.
	// It is the largest value m such that the address of v is always zero mod m.
	// It is the same as the value returned by reflect.TypeOf(x).Align().
	// As a special case, if a variable s is of struct type and f is a field
	// within that struct, then Alignof(s.f) will return the required alignment
	// of a field of that type within a struct. This case is the same as the
	// value returned by reflect.TypeOf(s.f).FieldAlign().
	b := true
	fmt.Println(unsafe.Alignof(b))

	fmt.Println()

	var s struct {
		f float64
		i int
		b bool
	}

	// Offsetof returns the offset within the struct of the field
	// represented by x, which must be of the form structValue.field.
	// In other words, it returns the number of bytes between the start
	// of the struct and the start of the field.
	fmt.Println(unsafe.Offsetof(s.b))

	fmt.Println()

	// Sizeof takes an expression x of any type and returns the size in bytes
	// of a hypothetical variable v as if v was declared via var v = x.
	// The size does not include any memory possibly referenced by x.
	// For instance, if x is a slice, Sizeof returns the size of the slice
	// descriptor, not the size of the memory referenced by the slice.
	fmt.Println(unsafe.Sizeof(s))

	fmt.Println()

	// Pointer arithmatic using unsafe.Pointer

	// Create a new integer array intArray
	intArray := [...]int{1, 2}

	// Print out intArray's elements
	fmt.Printf("intArray: %v\n", intArray)

	// Create and set intPtr to the first element's address in intArray
	intPtr := &intArray[0]

	// Print out the integer stored at the address pointed to by intPtr
	fmt.Printf("intPtr=%p, *intPtr=%d.\n", intPtr, *intPtr)

	// Move to the next element's address after the element at the
	// current position stored in intPtr
	//
	// Pointer represents a pointer to an arbitrary type.
	addressHolder := uintptr(unsafe.Pointer(intPtr)) + unsafe.Sizeof(intArray[0])

	// Set intPtr to the next addressHolder
	intPtr = (*int)(unsafe.Pointer(addressHolder))

	// Print out the integer stored at the address pointed to by intPtr
	// after reassigning intPtr to addressHolder
	fmt.Printf("intPtr=%p, *intPtr=%d.\n", intPtr, *intPtr)
}

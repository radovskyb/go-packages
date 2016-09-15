// Package subtle implements functions that are often
// useful in cryptographic code but require careful
// thought to use correctly.
package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	// ConstantTimeByteEq:

	X := uint8(1)
	Y := uint8(2)

	// ConstantTimeByteEq returns 1 if x == y and 0 otherwise.
	n := subtle.ConstantTimeByteEq(X, Y) // X != Y, n = 0
	fmt.Println(n)                       // 0

	X++ // Now equals 2

	n = subtle.ConstantTimeByteEq(X, Y) // X == Y, n = 1
	fmt.Println(n)                      // 1

	fmt.Println()

	// ConstantTimeCompare:

	A := []byte("Hello")
	B := []byte("World")
	C := []byte("Hello")

	// ConstantTimeCompare returns 1 if the two slices, x
	// and y, have equal contents. The time taken is a function
	// of the length of the slices and is independent of the contents.
	n = subtle.ConstantTimeCompare(A, B) // A != B, n = 0
	fmt.Println(n)                       // 0

	n = subtle.ConstantTimeCompare(A, C) // A == C, n = 1
	fmt.Println(n)                       // 1

	fmt.Println()

	// ConstantTimeCopy:

	H := []byte("Hello")
	I := make([]byte, len(H))

	// ConstantTimeCopy copies the contents of y into x (a slice of equal length)
	// if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v
	// takes any other value.
	subtle.ConstantTimeCopy(0, I, H) // v = 0 - Will NOT copy I to H
	fmt.Printf("I: `%s`\n", I)       // I: ``

	subtle.ConstantTimeCopy(1, I, H) // v = 1 - Will copy I to H
	fmt.Printf("I: `%s`\n", I)       // I: `Hello`

	fmt.Println()

	// ConstantTimeEq:

	L := int32(1)
	M := int32(2)
	N := int32(1)

	// ConstantTimeEq returns 1 if x == y and 0 otherwise.
	n = subtle.ConstantTimeEq(L, M)
	fmt.Println(n) // 0

	n = subtle.ConstantTimeEq(L, N)
	fmt.Println(n) // 1

	fmt.Println()

	// ConstantTimeLessOrEq:

	O := 3
	P := 4
	Q := 1

	// ConstantTimeLessOrEq returns 1 if x <= y and 0 otherwise.
	// Its behavior is undefined if x or y are negative or > 2**31 - 1.
	n = subtle.ConstantTimeLessOrEq(O, P)
	fmt.Println(n) // 1

	n = subtle.ConstantTimeLessOrEq(O, Q)
	fmt.Println(n) // 0

	fmt.Println()

	// ConstantTimeSelect:

	R := 3
	S := 4

	// ConstantTimeSelect returns x if v is 1 and y if v is 0.
	// Its behavior is undefined if v takes any other value.
	n = subtle.ConstantTimeSelect(1, R, S) // v = 1
	fmt.Println(n)                         // Prints `R`, 3

	n = subtle.ConstantTimeSelect(0, R, S) // v = 0
	fmt.Println(n)                         // Prints `S`, 4
}

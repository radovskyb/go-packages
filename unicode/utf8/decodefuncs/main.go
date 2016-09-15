package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		// DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and
		// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
		// the encoding is invalid, it returns (RuneError, 1). Both are impossible
		// results for correct, non-empty UTF-8.
		//
		// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
		// out of range, or is not the shortest possible UTF-8 encoding for the
		// value. No other validation is performed.
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[:len(b)-size]
	}

	fmt.Println()

	str := "Hello, 世界"
	for len(str) > 0 {
		// DecodeLastRuneInString is like DecodeLastRune but its input is a string. If
		// s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid,
		// it returns (RuneError, 1). Both are impossible results for correct,
		// non-empty UTF-8.
		//
		// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
		// out of range, or is not the shortest possible UTF-8 encoding for the
		// value. No other validation is performed.
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[:len(str)-size]
	}

	fmt.Println()

	b = []byte("Hello, 世界")
	for len(b) > 0 {
		// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and
		// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
		// the encoding is invalid, it returns (RuneError, 1). Both are impossible
		// results for correct, non-empty UTF-8.
		//
		// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
		// out of range, or is not the shortest possible UTF-8 encoding for the
		// value. No other validation is performed.
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
	}

	fmt.Println()

	str = "Hello, 世界"
	for len(str) > 0 {
		// DecodeRuneInString is like DecodeRune but its input is a string. If s is
		// empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it
		// returns (RuneError, 1). Both are impossible results for correct, non-empty
		// UTF-8.
		//
		// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
		// out of range, or is not the shortest possible UTF-8 encoding for the
		// value. No other validation is performed.
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
}

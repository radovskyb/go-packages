package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

const src = `
// This line is a comment from scanned go code
if a < 10 {
	fmt.Println("a is less than 10")
}`

func main() {
	// Create a new scanner.Scanner object
	//
	// A Scanner implements reading of Unicode characters
	// and tokens from an io.Reader.
	var s scanner.Scanner

	// Initialize the scanner with a new strings
	// reader that reads from `src`
	//
	// Init initializes a Scanner with a new source and returns s.
	// Error is set to nil, ErrorCount is set to 0, Mode is set to GoTokens,
	// and Whitespace is set to GoWhitespace.
	s.Init(strings.NewReader(src))

	// To allow scanner to also scan go comments as tokens, set `s.Mode` to
	// the same scanner mode that is set in s.Init which initializes `s.Mode` to
	// `scanner.GoTokens`. However remove one field `scanner.SkipComments` from the
	// set of fields taken from `scanner.GoTokens`. This is the field responsible
	// for skipping comments as tokens so without it, comments will be tokenized.
	//
	// s.Mode = scanner.ScanIdents | scanner.ScanFloats |
	// 	scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings |
	// 	scanner.ScanComments

	// Create a new rune variable `tok` to hold each rune scanned
	// in from scanner `s`
	var tok rune

	// Before scanning the first token into `s`, check if the position where
	// `s` is currently looking, is valid, as well as the current position
	// of `s` as a string. If the position is not valid it will print `???`
	//
	// IsValid reports whether the position is valid.
	fmt.Printf("IsValid & String before calling Scan: %t - %s\n\n",
		s.IsValid(), s.String())

	// Scan until `tok` is a scanner.EOF token
	for tok != scanner.EOF {
		// Set `tok` to the next scanned token from scanner `s`
		//
		// Scan reads the next token or Unicode character from source and returns it.
		// It only recognizes tokens t for which the respective Mode bit
		// (1<<-t) is set.
		// It returns EOF at the end of the source. It reports scanner errors (read
		// and token errors) by calling s.Error, if not nil; otherwise it prints
		// an error message to os.Stderr.
		tok = s.Scan()

		if tok != scanner.EOF {
			// Print the line number and position immediately after the character
			// or token returned by the last call to `s.Scan()` as well as print
			// the most recently valid scanned token after calling `s.Scan()`.
			//
			// Lastly print the returned string from calling scanner.TokenString
			// with `tok` as it's parameter, which returns a printable string
			// for a token or Unicode character, then using `s.Peek()` print out the
			// next unicode character after the current token's position, but don't
			// advance the scanner's position when peeking at the next character
			//
			// Pos returns the position of the character immediately after
			// the character or token returned by the last call to Next or Scan.
			//
			// TokenText returns the string corresponding to the most
			// recently scanned token valid after calling `s.Scan()`.
			//
			// TokenString returns a printable string for a token
			// or Unicode character.
			//
			// Peek returns the next Unicode character in the source without advancing
			// the scanner. It returns EOF if the scanner's position is at the last
			// character of the source.
			fmt.Printf("At position %v: %v - TokenString: %v, Peek: %v\n",
				s.Pos(), s.TokenText(), scanner.TokenString(tok),
				scanner.TokenString(s.Peek()))

			if s.Peek() == '\n' {
				fmt.Println("Newline")
			}
		} else {
			// Next reads and returns the next Unicode character.
			// It returns EOF at the end of the source. It reports
			// a read error by calling s.Error, if not nil; otherwise
			// it prints an error message to os.Stderr. Next does not
			// update the Scanner's Position field; use Pos() to
			// get the current position.
			//
			// Unlike peek, next does advance the scanners current position
			fmt.Println(scanner.TokenString(s.Next()))
		}
	}

	// Now that scanner has scanned until scanner.EOF, once again check if
	// the current position of `s` is valid and it's position's as a string
	fmt.Printf("\nIsValid & String after calling Scan: %t - %s\n",
		s.IsValid(), s.String())

}

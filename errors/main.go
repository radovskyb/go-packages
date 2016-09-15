package main

import (
	"errors"
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

// MyError now implicitly implememts the error interface
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// OopsieDaisies returns an error, in this case a MyError error since
// it implements from the error interface. If it didn't, returning
// the &MyError object would not be valid to return as it would
// not be classified as an error without implementing the error interface
func OopsieDaisies() error {
	return &MyError{
		time.Now(),
		"There has been a woopsie dasies error.",
	}
}

// Create a new error `slurpyError`
//
// New returns an error that formats as the given text.
var slurpyError = errors.New("Error: No slurpy found")

func Slurpy(hasSlurpy bool) error {
	if hasSlurpy {
		return nil
	}
	return slurpyError
}

func ErrorReturningFunction() error {
	// Errorf formats according to a format specifier and returns the string
	// as a value that satisfies error.
	return fmt.Errorf("Error: ErrorReturningFunction")
}

func main() {
	if err := OopsieDaisies(); err != nil {
		fmt.Println(err)
	}

	if err := Slurpy(false); err != nil {
		fmt.Println(err)
	}

	if err := ErrorReturningFunction(); err != nil {
		fmt.Println(err)
	}
}

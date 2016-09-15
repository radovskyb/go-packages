package main

import "fmt"

func Error() error {
	// Errorf formats according to a format specifier and returns the string
	// as a value that satisfies error.
	return fmt.Errorf("This is an error formatted and returned with %s", "fmt.Errorf")
}

func main() {
	if err := Error(); err != nil {
		fmt.Println(err)
	}
}

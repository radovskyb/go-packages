package main

import "fmt"

type MyObject struct {
	MyField int
}

// MyObject now implements fmt's GoStringer interface since it now
// has a GoString() method attached to it
//
// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value. The GoString method
// is used to print values passed as an operand to a %#v format.
func (mo MyObject) GoString() string {
	return fmt.Sprintf("MyObject{MyField: %d}", mo.MyField)
}

func main() {
	// Create a new MyObject object
	mo := &MyObject{25}

	// Print out mo's GoString() method since it is being
	// passed to fmt.Printf as an operand to a %#v format.
	fmt.Printf("%#v\n", mo)
}

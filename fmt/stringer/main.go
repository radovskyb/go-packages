package main

import "fmt"

type MyObject struct {
	MyField int
}

// MyObject now implements fmt's Stringer interface since it now
// has a String() method attached to it
//
// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value. The String method
// is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (mo MyObject) String() string {
	return fmt.Sprintf("MyObject{MyField: %d}", mo.MyField)
}

func main() {
	// Create a new MyObject object
	mo := &MyObject{25}

	// Print out mo's String() method since it is being printed
	// using fmt.Println()
	fmt.Println(mo)
}

package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	fmt.Println(xml.Header)
}

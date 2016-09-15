package main

import (
	"fmt"
	"strings"
)

func main() {
	// Index returns the index of the first instance of sep
	// in s, or -1 if sep is not present in s.
	fmt.Println(strings.Index("Hello, World!", "o")) // 4
	fmt.Println(strings.Index("Hello, World!", "z")) // -1
}

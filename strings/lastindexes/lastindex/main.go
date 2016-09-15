package main

import (
	"fmt"
	"strings"
)

func main() {
	// LastIndex returns the index of the last instance of sep
	// in s, or -1 if sep is not present in s.
	fmt.Println(strings.LastIndex("Hello, World!", "o")) // 8
	fmt.Println(strings.LastIndex("Hello, World!", "z")) // -1
}

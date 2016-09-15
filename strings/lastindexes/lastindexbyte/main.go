package main

import (
	"fmt"
	"strings"
)

func main() {
	// LastIndexByte returns the index of the last instance
	// of c in s, or -1 if c is not present in s.
	fmt.Println(strings.LastIndexByte("Hello, World!", byte('o'))) // 8
	fmt.Println(strings.LastIndexByte("Hello, World!", byte('z'))) // -1
}

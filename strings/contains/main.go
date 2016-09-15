package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello, World!"
	b := "Hello"
	c := "Nope"

	// Contains reports whether substr is within s.
	if strings.Contains(a, b) {
		fmt.Printf("'%s' contains '%s'\n", a, b)
	}

	if !strings.Contains(a, c) {
		fmt.Printf("'%s' doesn't contain '%s'\n", a, c)
	}

	fmt.Println(strings.Contains("radovsky", "radov"))
	fmt.Println(strings.Contains("radovsky", "rados"))
}

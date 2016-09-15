package main

import (
	"fmt"
	"strings"
)

func main() {
	// ContainsAny reports whether any Unicode code points in
	// chars are within s.
	//
	// false since there is no `i` in `team`
	fmt.Println(strings.ContainsAny("team", "i"))

	// true since `a` is in failure
	fmt.Println(strings.ContainsAny("failure", "abc"))
}

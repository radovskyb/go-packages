package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	t1 := time.Now().Add(time.Millisecond)
	t2 := t0

	// Equal reports whether t and u represent the same time instant.
	// Two times can be equal even if they are in different locations.
	// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
	// This comparison is different from using t == u, which also compares
	// the locations.
	fmt.Println(t0.Equal(t1)) // false
	fmt.Println(t0.Equal(t2)) // true
}

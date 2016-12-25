// The delete built-in function deletes the element with the
// specified key (m[key]) from the map. If m is nil or there
// is no such element, delete is a no-op.
// func delete(m map[Type]Type1, key Type) <- Function signature
package main

import "fmt"

func main() {
	// Create a map of [string]int
	m := make(map[string]int)

	// Add three separate elements into our map m
	m["first"] = 1
	m["second"] = 2
	m["third"] = 3

	// Print out the map before the deletion
	fmt.Println(m)

	// Print out the maps length before the deletion
	fmt.Println("length before deletion:", len(m))

	// Delete m["third"] from the map m
	delete(m, "third")

	// Print out the map after the deletion
	fmt.Printf("\n%v\n", m)

	// Print out the maps length after the deletion
	fmt.Println("length after deletion:", len(m))
}

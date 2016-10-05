// The cap built-in function returns the capacity of v,
// according to its type
// The len built-in function returns the length of v,
// according to its type:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// String len: the number of bytes in v
	fmt.Println("Strings:")

	fmt.Println("\nString One:")
	str := "Hello, World!"
	fmt.Println(str)
	fmt.Println("len:", len(str)) // 13

	fmt.Println()

	// Array cap: the number of elements in v (same as len(v)).
	// Array len: the number of elements in v.
	fmt.Println("Arrays:")

	fmt.Println("\nArray One:")
	var arr_one [3]int
	fmt.Println(arr_one)
	fmt.Println("cap:", cap(arr_one)) // 3
	fmt.Println("len:", len(arr_one)) // 3

	fmt.Println()

	// Pointer to array cap: the number of elements in *v (same as len(v)).
	// Pointer to array len: the number of elements in *v (even if v is nil).
	fmt.Println("Pointers:")
	ptr := new([]int)
	*ptr = []int{1, 2}
	fmt.Println(ptr)
	fmt.Println("cap:", cap(*ptr)) // 2
	fmt.Println("len:", len(*ptr)) // 2

	fmt.Println()

	// Slice cap: the maximum length the slice can reach when resliced;
	// Slice & Map len: the number of elements in v
	fmt.Println("Slices:")

	fmt.Println("\nSlice One:")
	slice_one := []int{1, 2, 3}
	fmt.Println(slice_one)
	fmt.Println("cap:", cap(slice_one)) // 3
	fmt.Println("len:", len(slice_one)) // 3

	fmt.Println("\nSlice Two:")
	slice_two := [10]int{1, 2, 3}
	fmt.Println(slice_two)
	fmt.Println("cap:", cap(slice_two)) // 10
	fmt.Println("len:", len(slice_two)) // 10

	fmt.Println("\nSlice Three:")
	slice_three := make([]int, 0, 5)
	fmt.Println(slice_three)
	fmt.Println("cap:", cap(slice_three)) // 5
	fmt.Println("len:", len(slice_three)) // 0

	fmt.Println("\nSlice Four:")
	slice_four := slice_three[:2]
	fmt.Println(slice_four)
	fmt.Println("cap:", cap(slice_four)) // 5
	fmt.Println("len:", len(slice_four)) // 2

	fmt.Println("\nSlice Five:")
	slice_five := slice_four[2:5]
	fmt.Println(slice_five)
	fmt.Println("cap:", cap(slice_five)) // 3
	fmt.Println("len:", len(slice_five)) // 3

	fmt.Println()

	// Channel cap: the channel buffer capacity, in units of elements;
	// Channel len: the number of elements queued (unread) in the channel buffer
	fmt.Println("Channels:")

	fmt.Println("\nChannel Two:")
	channel_one := make(chan int, 5)
	fmt.Println(reflect.TypeOf(channel_one))
	fmt.Println("cap:", cap(channel_one)) // 5
	fmt.Println("len:", len(channel_one)) // 0
	close(channel_one)

	fmt.Println("\nChannel Two:")
	channel_two := make(chan int, 5)
	channel_two <- 1
	channel_two <- 2
	fmt.Println(reflect.TypeOf(channel_two))
	fmt.Println("cap:", cap(channel_two)) // 5
	fmt.Println("len:", len(channel_two)) // 2
	close(channel_two)
}

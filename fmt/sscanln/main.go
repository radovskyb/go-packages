package main

import "fmt"

func main() {
	var first, second, third string

	fmt.Sscanln("First Second\nThird", &first, &second, &third)

	fmt.Printf("First: %s\nSecond %s\n", first, second)

	if third == "" {
		fmt.Println("`third` wasn't scanned into since after scanning " +
			"into second, there was a \\n character.")
	}
}

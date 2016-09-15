package main

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	fmt.Printf("Year: %v\nMonth: %v\nDay: %v\n", year, month, day)
}

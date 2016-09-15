package main

import (
	"fmt"
	"sort"
)

type Grams int

func (g Grams) String() string {
	return fmt.Sprintf("%dg", int(g))
}

type Food struct {
	Name   string
	Weight Grams
}

type Foods []*Food

func (f Foods) Len() int {
	return len(f)
}

func (f Foods) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type ByName struct {
	Foods
}

func (n ByName) Less(i, j int) bool {
	return n.Foods[i].Name < n.Foods[j].Name
}

type ByWeight struct {
	Foods
}

func (w ByWeight) Less(i, j int) bool {
	return w.Foods[i].Weight < w.Foods[j].Weight
}

var foods = []*Food{
	{"Lollypop", 100},
	{"Chocolate Bar", 250},
	{"Apple", 125},
	{"Watermelon", 500},
}

func main() {
	fmt.Println("Unsorted:")
	printFoods(foods)

	sort.Sort(ByName{foods})
	fmt.Println("Sorted By Name:")
	printFoods(foods)

	sort.Sort(ByWeight{foods})
	fmt.Println("Sorted By Weight:")
	printFoods(foods)
}

func printFoods(foods []*Food) {
	for _, food := range foods {
		fmt.Printf("%s: %d\n", food.Name, food.Weight)
	}
	fmt.Println()
}

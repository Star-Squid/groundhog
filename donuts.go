package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out the inventory
	fmt.Println(donuts)

	secondChoice, status := donuts["chocolate"]

	if status {
		fmt.Println(secondChoice)
	} else {
		fmt.Println("We do not sell that donut!")
	}
}

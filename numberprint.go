package main

import "fmt"

func printItAll() {

	source := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num1 := range source {
		for _, num2 := range source {

			if num1 < num2 {
				for _, num3 := range source {
					if num2 < num3 {
						// fmt.Print(num1, num2, num3, ", ")
						fmt.Printf("%d%d%d, ", num1, num2, num3)
					}
				}
			}
		}
	}
}

func main() {
	printItAll()
}

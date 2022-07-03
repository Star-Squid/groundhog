package main

import (
	"fmt"
	"quad/piscine"
)

func main() {
	//Edit these values to change number of test cases
	xMin := -1
	xMax := 5
	yMin := -2
	yMax := 7
	//Run test cases
	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			//Print result of test case
			fmt.Printf("TEST CASE\nx: %d\ty: %d\n", i, j)
			fmt.Printf("\nOutput:\n")
			piscine.QuadE(i, j)
			fmt.Printf("\n")
		}
	}
}

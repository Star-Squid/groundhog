package main

import "fmt"

func calculateProportion(num1, num2, num3 float64) float64 {
	var divider = num1 / num2
	var result = num3 / divider
	return result
}

func main() {

	fmt.Println("Example:")
	fmt.Println("a -- b")
	fmt.Println("c -- X1")

	var a, b, c, x float64

	fmt.Println("What's a?")
	fmt.Scan(&a)
	fmt.Println("What's b?")
	fmt.Scan(&b)
	fmt.Println("What's c?")
	fmt.Scan(&c)

	x = calculateProportion(a, b, c)

	fmt.Println("The value of x is", x)
}

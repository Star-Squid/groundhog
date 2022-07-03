package main

import "github.com/01-edu/z01"

func main() {
	source := "0123456789"
	runes := []rune(source)

	for i := 0; i < len(source); i++ {
		z01.PrintRune(runes[i])
	}

	z01.PrintRune('\n')
}

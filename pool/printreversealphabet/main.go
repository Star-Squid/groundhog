package main

import "github.com/01-edu/z01"

func main() {
	source := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(source)

	for i := len(source) - 1; i >= 0; i-- {
		z01.PrintRune(runes[i])
	}

	z01.PrintRune('\n')
}

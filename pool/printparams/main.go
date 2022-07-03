package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	allArgs := os.Args[1:]
	//{choumi is the best cat}

	for i := 0; i < len(allArgs); i++ {

		runes := []rune(allArgs[i])
		for i := 0; i < len(runes); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}

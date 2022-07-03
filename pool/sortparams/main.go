package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	allArgs := os.Args[1:]

	if len(allArgs) > 0 {
		// double loop sorts the array
		for i := 0; i < len(allArgs)-1; i++ {
			for i := 0; i < len(allArgs)-1; i++ {
				if allArgs[i] > allArgs[i+1] {
					allArgs[i], allArgs[i+1] = allArgs[i+1], allArgs[i]
				}
			}
		}

		for i := 0; i < len(allArgs); i++ {
			// convert each arg to runes, print each rune
			runes := []rune(allArgs[i])

			for i := 0; i < len(runes); i++ {
				z01.PrintRune(runes[i])
			}
			z01.PrintRune('\n')
		}
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	runes := []rune(programName)

	// find last slash position
	lastSlash := 0

	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] == '/' {
			lastSlash = i
		}
	}

	// print after the slash
	for i := lastSlash + 1; i <= len(runes)-1; i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}

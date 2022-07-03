package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	runes := []int{}

	if n <= 0 {
		z01.PrintRune('0')
	}

	for i := n; i > 0; i /= 10 {
		num := i % 10
		runes = append(runes, num)
	}

	if len(runes) > 0 {
		for i := 0; i < len(runes)-1; i++ {
			for i := 0; i < len(runes)-1; i++ {
				if runes[i] > runes[i+1] {
					runes[i], runes[i+1] = runes[i+1], runes[i]
				}
			}
		}
	}

	for i := 0; i < len(runes); i++ {
		runeFromNum := rune(48 + runes[i])
		z01.PrintRune(runeFromNum)
	}
}

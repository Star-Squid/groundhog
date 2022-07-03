package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	allArgs := os.Args[1:]

	if len(allArgs) == 1 {
		s := allArgs[0]
		number, _ := strconv.Atoi(s)
		if CheckPowerOfTwo(number) == 0 {
			printTrue()
		} else {
			printFalse()
		}
	}
}

func CheckPowerOfTwo(n int) int {
	// added one corner case if n is zero it will also consider as power 2
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func printTrue() {
	z01.PrintRune('t')
	z01.PrintRune('r')
	z01.PrintRune('u')
	z01.PrintRune('e')
	z01.PrintRune('\n')
}

func printFalse() {
	z01.PrintRune('f')
	z01.PrintRune('a')
	z01.PrintRune('l')
	z01.PrintRune('s')
	z01.PrintRune('e')
	z01.PrintRune('\n')
}

package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if isBaseValid(base) {
		for _, num := range base {
			z01.PrintRune(num)
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
		z01.PrintRune('\n')
	}
}

func isBaseValid(b string) bool {
	if Atoi(b) <= 0 {
		return true
	} else {
		return false
	}
}

// Write a function that prints an int in a string base passed as parameters.

// If the base is not valid, the function prints NV (Not Valid):

// Validity rules for a base :

//     A base must contain at least 2 characters.
//     Each character of a base must be unique.
//     A base should not contain + or - characters.

// The function has to manage negative numbers. (as shown in the example)

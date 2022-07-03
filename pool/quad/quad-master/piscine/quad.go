package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x int, y int) {
	//If x or y <= 0, immediately return
	if x <= 0 || y <= 0 {
		return
	}
	//Declare and initialise a rune slice with the relevant symbols
	runes := []rune{'o', '|', '-', ' '}
	//xMod and yMod are the moduli for the columns and rows respectively (0-indexed - hence x-1, y-1 assignment)
	//Use these to check if we are in left/right column and top/bottom row.
	//E.g. Column: (0 % xMod) = 0 and (xMod % xMod) = 0. The same holds for Rows.
	xMod := x - 1
	yMod := y - 1
	//Special case: only 1 column or row. xMod, yMod cannot be 0, as division by 0 is undefined.
	if x == 1 {
		xMod = 1
	}
	if y == 1 {
		yMod = 1
	}
	//Traverse the rectangle, row by row
	for row := 0; row < y; row++ {
		//Traverse each row, column by column
		for col := 0; col < x; col++ {
			//Prepare to shift the desired index of the rune slice if not left/right column
			colShift := Shift(col, xMod)
			//Similar as above, except for top/bottom row
			rowShift := Shift(row, yMod)
			//Set desired index to Decimal Representation of Possible 4 Binary Values
			//I.E. Corner: 		    2(0) + 0 = 0 DEC = 00 BIN
			//	   Top/Bottom Only:	2(0) + 1 = 1 DEC = 01 BIN
			//	   Left/Right Only: 2(1) + 0 = 2 DEC = 10 BIN
			//	   Neither:			2(1) + 1 = 3 DEC = 11 BIN
			index := 2*colShift + rowShift
			//Print out desired rune
			z01.PrintRune(runes[index])
		}
		//Print newline character before moving onto next row
		z01.PrintRune('\n')
	}
}

func QuadB(x int, y int) {
	//If x or y <= 0, immediately return
	if x <= 0 || y <= 0 {
		return
	}
	//Declare and initialise a rune slice with the relevant symbols.
	//Note that we only have one corner rune in this slice, and will apply
	//an ASCII conversion if necessary.
	runes := []rune{'/', '*', ' '}
	//xMod and yMod are the moduli for the columns and rows respectively (0-indexed - hence x-1, y-1 assignment)
	//Use these to check if we are in left/right column and top/bottom row.
	//E.g. Column: (0 % xMod) = 0 and (xMod % xMod) = 0. The same holds for Rows.
	xMod := x - 1
	yMod := y - 1
	//Special case: only 1 column or row. xMod, yMod cannot be 0, as division by 0 is undefined.
	if x == 1 {
		xMod = 1
	}
	if y == 1 {
		yMod = 1
	}
	//Traverse the rectangle, row by row
	for row := 0; row < y; row++ {
		//Traverse each row, column by column
		for col := 0; col < x; col++ {
			//Prepare to shift the desired index of the rune slice if not left/right column
			colShift := Shift(col, xMod)
			//Similar as above, except for top/bottom row
			rowShift := Shift(row, yMod)

			index := colShift + rowShift
			// This sets our desired index to
			// Corner: 0 + 0 = 0
			// Edge: 0 + 1 or 1 + 0 = 1
			// Principal Sub-Matrix (Centre): 1 + 1 = 2
			//
			// Finally, we have to check whether the corners
			// are in the top-left, top-right, bottom-left or bottom-right of the grid
			output := runes[index]
			if index == 0 {
				// colDiv = 0 for left column, colDiv = 1 for right column
				colDiv := col / xMod
				// rowDiv = 0 for top row, rowDiv = 1 for bottom row
				rowDiv := row / yMod
				// If the absolute difference between colDiv and rowDiv = 1,
				// we know that we are either in the top-right or bottom-left corner
				// and thus apply the corner conversion to the rune.
				if AbsDifference(colDiv, rowDiv) == 1 {
					cornerConversion := '\\' - '/'
					output += cornerConversion
				}
			}
			z01.PrintRune(output)
		}
		//Print newline character before moving onto next row
		z01.PrintRune('\n')
	}
}

func QuadC(x int, y int) {
	//If x or y <= 0, immediately return
	if x <= 0 || y <= 0 {
		return
	}
	//Declare and initialise a rune slice with the relevant symbols.
	//Note that we only have one corner rune in this slice, and will apply
	//an ASCII conversion if necessary.
	runes := []rune{'A', 'B', ' '}
	//xMod and yMod are the moduli for the columns and rows respectively (0-indexed - hence x-1, y-1 assignment)
	//Use these to check if we are in left/right column and top/bottom row.
	//E.g. Column: (0 % xMod) = 0 and (xMod % xMod) = 0. The same holds for Rows.
	xMod := x - 1
	yMod := y - 1
	//Special case: only 1 column or row. xMod, yMod cannot be 0, as division by 0 is undefined.
	if x == 1 {
		xMod = 1
	}
	if y == 1 {
		yMod = 1
	}
	//Traverse the rectangle, row by row
	for row := 0; row < y; row++ {
		//Traverse each row, column by column
		for col := 0; col < x; col++ {
			//Prepare to shift the desired index of the rune slice if not left/right column
			colShift := Shift(col, xMod)
			//Similar as above, except for top/bottom row
			rowShift := Shift(row, yMod)

			index := colShift + rowShift
			// This sets our desired index to
			// Corner: 0 + 0 = 0
			// Edge: 0 + 1 or 1 + 0 = 1
			// Principal Sub-Matrix (Centre): 1 + 1 = 2
			//
			// Finally, we have to check whether the corners
			// are in the top row or bottom row of the grid, and if they
			// are in the bottom row, apply the appropriate rune conversion
			output := runes[index]
			if index == 0 {
				// rowDiv = 0 for top row, rowDiv = 1 for bottom row
				rowDiv := row / yMod
				if rowDiv == 1 {
					output += 'C' - 'A'
				}
			}
			z01.PrintRune(output)
		}
		//Print newline character before moving onto next row
		z01.PrintRune('\n')
	}
}
func QuadD(x int, y int) {
	//If x or y <= 0, immediately return
	if x <= 0 || y <= 0 {
		return
	}
	//Declare and initialise a rune slice with the relevant symbols.
	//Note that we only have one corner rune in this slice, and will apply
	//an ASCII conversion if necessary.
	runes := []rune{'A', 'B', ' '}
	//xMod and yMod are the moduli for the columns and rows respectively (0-indexed - hence x-1, y-1 assignment)
	//Use these to check if we are in left/right column and top/bottom row.
	//E.g. Column: (0 % xMod) = 0 and (xMod % xMod) = 0. The same holds for Rows.
	xMod := x - 1
	yMod := y - 1
	//Special case: only 1 column or row. xMod, yMod cannot be 0, as division by 0 is undefined.
	if x == 1 {
		xMod = 1
	}
	if y == 1 {
		yMod = 1
	}
	//Traverse the rectangle, row by row
	for row := 0; row < y; row++ {
		//Traverse each row, column by column
		for col := 0; col < x; col++ {
			//Prepare to shift the desired index of the rune slice if not left/right column
			colShift := Shift(col, xMod)
			//Similar as above, except for top/bottom row
			rowShift := Shift(row, yMod)

			index := colShift + rowShift
			// This sets our desired index to
			// Corner: 0 + 0 = 0
			// Edge: 0 + 1 or 1 + 0 = 1
			// Principal Sub-Matrix (Centre): 1 + 1 = 2
			//
			// Finally, we have to check whether the corners
			// are in the left column or right of the grid, and if they
			// are in the right column, apply the appropriate rune conversion
			output := runes[index]
			if index == 0 {
				// colDiv = 0 for left column, colDiv = 1 for right column
				colDiv := col / xMod
				if colDiv == 1 {
					output += 'C' - 'A'
				}
			}
			z01.PrintRune(output)
		}
		//Print newline character before moving onto next row
		z01.PrintRune('\n')
	}
}

func QuadE(x int, y int) {
	//If x or y <= 0, immediately return
	if x <= 0 || y <= 0 {
		return
	}
	//Declare and initialise a rune slice with the relevant symbols.
	//Note that we only have one corner rune in this slice, and will apply
	//an ASCII conversion if necessary.
	runes := []rune{'A', 'B', ' '}
	//xMod and yMod are the moduli for the columns and rows respectively (0-indexed - hence x-1, y-1 assignment)
	//Use these to check if we are in left/right column and top/bottom row.
	//E.g. Column: (0 % xMod) = 0 and (xMod % xMod) = 0. The same holds for Rows.
	xMod := x - 1
	yMod := y - 1
	//Special case: only 1 column or row. xMod, yMod cannot be 0, as division by 0 is undefined.
	if x == 1 {
		xMod = 1
	}
	if y == 1 {
		yMod = 1
	}
	//Traverse the rectangle, row by row
	for row := 0; row < y; row++ {
		//Traverse each row, column by column
		for col := 0; col < x; col++ {
			//Prepare to shift the desired index of the rune slice if not left/right column
			colShift := Shift(col, xMod)
			//Similar as above, except for top/bottom row
			rowShift := Shift(row, yMod)

			index := colShift + rowShift
			// This sets our desired index to
			// Corner: 0 + 0 = 0
			// Edge: 0 + 1 or 1 + 0 = 1
			// Principal Sub-Matrix (Centre): 1 + 1 = 2
			//
			// Finally, we have to check whether the corners
			// are in the top-left, top-right, bottom-left or bottom-right of the grid
			output := runes[index]
			if index == 0 {
				// colDiv = 0 for left column, colDiv = 1 for right column
				colDiv := col / xMod
				// rowDiv = 0 for top row, rowDiv = 1 for bottom row
				rowDiv := row / yMod
				// If the absolute difference between colDiv and rowDiv = 1,
				// we know that we are either in the top-right or bottom-left corner
				// and thus apply the corner conversion to the rune.
				if AbsDifference(colDiv, rowDiv) == 1 {
					cornerConversion := 'C' - 'A'
					output += cornerConversion
				}
			}
			z01.PrintRune(output)
		}
		//Print newline character before moving onto next row
		z01.PrintRune('\n')
	}
}
func AbsDifference(m int, n int) int {
	if m-n < 0 {
		return n - m
	}
	return m - n
}
func Shift(div int, mod int) int {
	if div%mod == 0 {
		return 0
	}
	return 1
}

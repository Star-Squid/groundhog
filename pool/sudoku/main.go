package main

import (
	"os"

	"github.com/01-edu/z01"

	"fmt"
)

func main() {
	args := os.Args[1:]

	//check if args are valid
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	//make the board
	board := make([][]rune, 9)

	for i := 0; i < len(board); i++ {
		for _, num := range args[i] {
			if num == '.' {
				board[i] = append(board[i], '0')
			} else if (num >= '1' || num <= '9') && num != '.' {
				board[i] = append(board[i], num)
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	sudoku(board, 0, 0) // recursively calls the possible function for each grid cell and then places the valid numbers

	for _, row := range board {
		for _, col := range row {
			if !(col >= '1' && col <= '9') {
				fmt.Println("Error")
				return
			}
			z01.PrintRune(col)
		}
		z01.PrintRune('\n')
	}
}

func possible(theboard [][]rune, num rune, row int, col int) bool {
	//is num on row or on column?
	for i := 0; i < 9; i++ {
		if theboard[row][i] == num || theboard[i][col] == num {
			return false
		}
	}

	//is num in box?
	startofrow := row - row%3
	startofcol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if theboard[startofrow+i][startofcol+j] == num {
				return false
			}
		}
	}

	return true

}

func sudoku(theboard [][]rune, row int, col int) bool {

	if row == 8 && col == 9 {
		return true
	}

	if col == 9 {
		row++
		col = 0
	}

	if theboard[row][col] >= '1' && theboard[row][col] <= '9' {
		return sudoku(theboard, row, col+1)
	}

	for i := '1'; i <= '9'; i++ {
		if possible(theboard, i, row, col) {
			theboard[row][col] = i
			if sudoku(theboard, row, col+1) {
				return true
			}
			theboard[row][col] = 0
		}
	}
	return false
}

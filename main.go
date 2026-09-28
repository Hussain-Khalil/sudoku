package main

import (
	"fmt"
	"os"

	SUDOKU "SUDOKU/Functions"
)

func main() {
	input := os.Args[1:]

	matrix := SUDOKU.OstoArray(input)
	if err := SUDOKU.CheckSudokuErrors(matrix); err != nil {
		fmt.Println("Error")
		return
	}
	if SUDOKU.SolveSudoku(matrix) {
		for _, row := range matrix {
			for _, v := range row {
				fmt.Print(v)
				fmt.Print(" ")
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("Error!")
	}
}

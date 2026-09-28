package SUDOKU

import (
	"strconv"
)

func OstoArray(s []string) [][]int {
	matrix := make([][]int, len(s))
	for i := range matrix {
		matrix[i] = make([]int, len(s[i]))
	}

	for i, str := range s {
		for j, char := range str {
			if char == '.' {
				matrix[i][j] = 0
			} else {
				num, _ := strconv.Atoi(string(char))
				matrix[i][j] = num
			}
		}
	}

	return matrix
}

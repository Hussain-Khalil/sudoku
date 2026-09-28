package SUDOKU

import "fmt"

func CheckSudokuErrors(matrix [][]int) error {
	if len(matrix) != 9 {
		return fmt.Errorf("Error")
	}
	for _, v := range matrix {
		if len(v) != 9 {
			return fmt.Errorf("Error")
		}
	}
	for i := 0; i < 9; i++ {
		rowSet := make(map[int]bool)
		colSet := make(map[int]bool)

		for j := 0; j < 9; j++ {
			if matrix[i][j] != 0 {
				if rowSet[matrix[i][j]] {
					return fmt.Errorf("Error")
				}
				rowSet[matrix[i][j]] = true
			}

			if matrix[j][i] != 0 {
				if colSet[matrix[j][i]] {
					return fmt.Errorf("Error")
				}
				colSet[matrix[j][i]] = true
			}
		}
	}

	for blockRow := 0; blockRow < 3; blockRow++ {
		for blockCol := 0; blockCol < 3; blockCol++ {
			blockSet := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := matrix[blockRow*3+i][blockCol*3+j]
					if val != 0 {
						if blockSet[val] {
							return fmt.Errorf("Error")
						}
						blockSet[val] = true
					}
				}
			}
		}
	}

	return nil
}

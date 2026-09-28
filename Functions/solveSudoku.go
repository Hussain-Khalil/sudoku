package SUDOKU



func SolveSudoku(matrix [][]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if matrix[r][c] == 0 {
				for possibleValue := 1; possibleValue <= 9; possibleValue++ {
					if IsPossibleValue(possibleValue, r, c, matrix) {
						matrix[r][c] = possibleValue
						if SolveSudoku(matrix) {
							return true
						}
						matrix[r][c] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

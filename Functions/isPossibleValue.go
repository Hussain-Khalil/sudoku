package SUDOKU

func IsPossibleValue(possibleValue, r, c int, matrix [][]int) bool {
	for i := 0; i < 9; i++ {
		if matrix[r][i] == possibleValue || matrix[i][c] == possibleValue {
			return false
		}
	}

	blockRowStart := (r / 3) * 3
	blockColStart := (c / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[blockRowStart+i][blockColStart+j] == possibleValue {
				return false
			}
		}
	}

	return true
}

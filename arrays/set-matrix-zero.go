package main

func MatrixZero(matrix [][]int) {
	rows := make([]int, len(matrix))
	cols := make([]int, len(matrix[0]))

	// mark the row and col with 1 to convert them into 0 later
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == 0 {
				rows[row] = 1
				cols[col] = 1
			}
		}
	}

	// update the matrix to zero on the marked index
	for row, rowValue := range rows {
		for col, colValue := range cols {
			if rowValue == 1 || colValue == 1 {
				matrix[row][col] = 0
			}
		}
	}
}

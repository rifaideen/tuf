package main

// Search in a row and column wise sorted matrix
func searchMatrixBS2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

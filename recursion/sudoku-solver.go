package main

func solveSudoku(board [][]byte) {
	solveSudokuRecursively(board)
}

func solveSudokuRecursively(board [][]byte) bool {
	// traverse the board and find the empty cell
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				// check if we can place number between 1 to 9 at the current cell
				for i := byte('1'); i <= byte('9'); i++ {
					if canPlace(board, row, col, i) {
						// set the cell with value i
						board[row][col] = i

						// check again with other empty cells
						if solveSudokuRecursively(board) {
							return true
						}

						// backtrack the board by removing the value added to the cell
						board[row][col] = '.'
					}
				}

				// since we could not place any numbers in the above loop, returning false would be appropriate
				return false
			}
		}
	}

	// we are able to fill and solve, hence true
	return true
}

func canPlace(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		// check row-wise and column-wise
		if board[row][i] == val || board[i][col] == val {
			return false
		}

		// check board-wise (current board 3x3)
		r := (3 * (row / 3)) + (i / 3)
		c := (3 * (col / 3)) + (i % 3)

		if board[r][c] == val {
			return false
		}
	}

	return true
}

package main

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && isExists(0, i, j, rows, cols, word, board) {
				return true
			}
		}
	}

	return false
}

func isExists(index, i, j, rows, cols int, word string, board [][]byte) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || j < 0 || rows == i || cols == j || board[i][j] != word[index] || board[i][j] == '!' {
		return false
	}

	char := board[i][j]
	board[i][j] = '!'

	top := isExists(index+1, i-1, j, rows, cols, word, board)
	right := isExists(index+1, i, j+1, rows, cols, word, board)
	bottom := isExists(index+1, i+1, j, rows, cols, word, board)
	left := isExists(index+1, i, j-1, rows, cols, word, board)

	board[i][j] = char

	return top || bottom || left || right
}

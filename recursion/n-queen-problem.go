package main

import "strings"

func solveNQueens(n int) [][]string {
	// create board with dots
	board := make([][]string, 0, n)

	for row := 0; row < n; row++ {
		cols := make([]string, n)

		for col := 0; col < n; col++ {
			cols[col] = "."
		}

		board = append(board, cols)
	}

	ans := [][]string{}

	solveNQueensRecursive(0, n, &board, &ans)

	return ans
}

func solveNQueensRecursive(col, n int, board *[][]string, ans *[][]string) {
	// the moment col crosses the entire columns (n), we have successfully placed queen
	if col == n {
		// create boards of size n
		boards := make([]string, 0, n)

		// convert each row of board array into string and append them on boards
		for _, row := range *board {
			boards = append(boards, strings.Join(row, ""))
		}

		// append the converted boards into the ans
		*ans = append(*ans, boards)

		return
	}

	for row := range n {
		if canPlaceQueen(row, col, n, *board) {
			(*board)[row][col] = "Q"
			solveNQueensRecursive(col+1, n, board, ans)

			// back-track and remove
			(*board)[row][col] = "."
		}
	}
}

func canPlaceQueen(row, col, n int, board [][]string) bool {
	// copy the row and col indices while preserving the provided row and col index
	r := row
	c := col

	// move up diagonally to the left
	for r >= 0 && c >= 0 {
		if board[r][c] == "Q" {
			return false
		}

		// reducing columns and rows to move up diagonally
		r--
		c--
	}

	// reset the row and col position
	r = row
	c = col

	// move left
	for c >= 0 {
		if board[r][c] == "Q" {
			return false
		}

		// move the col to the left
		c--
	}

	r = row
	c = col

	// move down diagonally to the left
	for r < n && c >= 0 {
		if board[r][c] == "Q" {
			return false
		}

		// move the row below and move the col to the left
		r++
		c--
	}

	return true
}

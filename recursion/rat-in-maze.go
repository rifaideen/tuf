package main

import "strings"

// D|L|R|U
func ratInMaze(board [][]int) []string {
	ans := []string{}

	if len(board) == 0 {
		return ans
	}

	n := len(board[0])

	ratInMazeRecursive(0, 0, n, board, &[]string{}, &ans)

	return ans
}

func ratInMazeRecursive(row, col, n int, board [][]int, directions *[]string, ans *[]string) {
	if row == n-1 && col == n-1 {
		path := strings.Join(*directions, "")
		*ans = append(*ans, path)

		return
	}

	if row < 0 || col < 0 {
		return
	}

	// can move down?
	if canMove(row+1, col, n, board) {
		*directions = append(*directions, "D")
		// temporarily block the path
		board[row][col] = 0
		ratInMazeRecursive(row+1, col, n, board, directions, ans)

		// unblock the path and backtrack
		board[row][col] = 1
		*directions = (*directions)[:len(*directions)-1]
	}

	// can move left?
	if canMove(row, col-1, n, board) {
		*directions = append(*directions, "L")
		// temporarily block the path
		board[row][col] = 0
		ratInMazeRecursive(row, col-1, n, board, directions, ans)

		// unblock the path and backtrack
		board[row][col] = 1
		*directions = (*directions)[:len(*directions)-1]
	}

	// can move right?
	if canMove(row, col+1, n, board) {
		*directions = append(*directions, "R")
		// temporarily block the path
		board[row][col] = 0
		ratInMazeRecursive(row, col+1, n, board, directions, ans)

		// unblock the path and backtrack
		board[row][col] = 1
		*directions = (*directions)[:len(*directions)-1]
	}

	// can move up?
	if canMove(row-1, col, n, board) {
		*directions = append(*directions, "U")
		// temporarily block the path
		board[row][col] = 0
		ratInMazeRecursive(row-1, col, n, board, directions, ans)

		// unblock the path and backtrack
		board[row][col] = 1
		*directions = (*directions)[:len(*directions)-1]
	}
}

func canMove(row, col, n int, board [][]int) bool {
	if row == n || col == n || row < 0 || col < 0 {
		return false
	}

	return board[row][col] == 1
}

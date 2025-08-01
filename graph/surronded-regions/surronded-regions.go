package main

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// Create a visited matrix to track which 'O's are connected to the border.
	// This helps us distinguish between:
	//   - 'O's that are on the border or connected to the border (should NOT be flipped)
	//   - 'O's that are completely surrounded by 'X's (should be flipped to 'X')
	visited := make([][]int, rows)

	for i := range visited {
		visited[i] = make([]int, cols)
	}

	// Traverse the top row (first row) and bottom row (last row)
	// For every 'O' on the border, it cannot be surrounded, so it and all 'O's connected to it
	// must remain as 'O'. We start DFS from these border 'O's to mark all connected 'O's as safe (visited).
	for col := 0; col < cols; col++ {
		// If top-left or top-right or any top-row cell is 'O' and not yet visited,
		// it's on the border → mark it and all its connected 'O' neighbors as safe (via DFS)
		if board[0][col] == 'O' && visited[0][col] == 0 {
			dfs(0, col, rows, cols, board, visited)
		}

		// Same for the bottom row: any 'O' here is on the border
		// → so it and its connected region must be preserved
		if board[rows-1][col] == 'O' && visited[rows-1][col] == 0 {
			dfs(rows-1, col, rows, cols, board, visited)
		}
	}

	// Now traverse the leftmost and rightmost columns (from top to bottom)
	// Any 'O' on the left/right edge is also on the border → cannot be surrounded
	for row := 0; row < rows; row++ {
		// Leftmost column: if 'O' and not visited, start DFS to mark entire connected component as safe
		if board[row][0] == 'O' && visited[row][0] == 0 {
			dfs(row, 0, rows, cols, board, visited)
		}

		// Rightmost column: same logic
		if board[row][cols-1] == 'O' && visited[row][cols-1] == 0 {
			dfs(row, cols-1, rows, cols, board, visited)
		}
	}

	// Final pass: go through every cell in the board
	// Any 'O' that was NOT visited during DFS is NOT connected to the border
	// → which means it is completely surrounded by 'X's → should be flipped to 'X'
	for row := range rows {
		for col := range cols {
			// If the cell is 'O' but was never reached by DFS (i.e., not connected to border),
			// then it's an "inner" O → flip it to 'X'
			if board[row][col] == 'O' && visited[row][col] == 0 {
				board[row][col] = 'X'
			}
		}
	}
}

// DFS function: explores all connected 'O's starting from (row, col)
// Marks them as visited to indicate they are safe (connected to the border)
func dfs(row, col, rows, cols int, board [][]byte, visited [][]int) {
	visited[row][col] = 1 // Mark current cell as visited (safe)

	// Directions: up, down, left, right
	delRow := []int{-1, 1, 0, 0} // up, down, no vertical for left/right
	delCol := []int{0, 0, -1, 1} // no horizontal for up/down, left, right

	// Explore all 4 adjacent cells
	for i := range 4 {
		r := row + delRow[i] // new row index
		c := col + delCol[i] // new col index

		// Check if the new position is within bounds
		if r >= 0 && r < rows && c >= 0 && c < cols {
			// If the neighboring cell is 'O' and hasn't been visited yet,
			// recursively explore it — because it's part of the same safe region
			if visited[r][c] == 0 && board[r][c] == 'O' {
				dfs(r, c, rows, cols, board, visited)
			}
		}
	}
}

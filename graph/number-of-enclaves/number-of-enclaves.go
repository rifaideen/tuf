package main

// Function to count the number of enclaves: 1s that are not connected to the boundary
// and thus cannot "escape" — they are completely surrounded by water (0s) or other land (1s)
// but not touching the border.
func numEnclaves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// Visited matrix to track which land cells (1s) are reachable from the boundary
	visited := make([][]int, rows)

	// Initialize each row of the visited matrix
	for i := range rows {
		visited[i] = make([]int, cols)
	}

	// Queue to store all boundary land cells (starting points for our DFS)
	queue := [][]int{}

	// Traverse the top and bottom rows, add any land cell (1) to the queue
	for col := range cols {
		if grid[0][col] == 1 {
			queue = append(queue, []int{0, col}) // Top row
		}

		if grid[rows-1][col] == 1 {
			queue = append(queue, []int{rows - 1, col}) // Bottom row
		}
	}

	// Traverse the leftmost and rightmost columns, add any land cell (1) to the queue
	for row := range rows {
		if grid[row][0] == 1 {
			queue = append(queue, []int{row, 0}) // Left column
		}

		if grid[row][cols-1] == 1 {
			queue = append(queue, []int{row, cols - 1}) // Right column
		}
	}

	// DFS function to mark all land cells connected to the boundary
	// Any land we can reach from the boundary is *not* an enclave
	var dfs func(row, col int)
	dfs = func(row, col int) {
		visited[row][col] = 1 // Mark current cell as visited

		// Directions: up, down, left, right
		delRows := []int{-1, 1, 0, 0}
		delCols := []int{0, 0, -1, 1}

		// Explore all 4 neighboring cells
		for i := range 4 {
			r := row + delRows[i]
			c := col + delCols[i]

			// Check if the neighbor is within bounds
			if r >= 0 && r < rows && c >= 0 && c < cols {
				// If it's unvisited land, continue DFS
				if visited[r][c] == 0 && grid[r][c] == 1 {
					dfs(r, c)
				}
			}
		}
	}

	// Perform DFS from every land cell on the boundary
	// This marks all land cells connected to the boundary as visited
	for _, node := range queue {
		dfs(node[0], node[1])
	}

	// Now, count all land cells (1s) that were *not* visited
	// These are the enclaves — land not connected to any boundary
	enclaves := 0
	for row := range rows {
		for col := range cols {
			if grid[row][col] == 1 && visited[row][col] == 0 {
				enclaves++
			}
		}
	}

	// Return the total number of enclaves
	return enclaves
}

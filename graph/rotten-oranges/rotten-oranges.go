package main

/*
INTUITION:
This is a multi-source BFS problem. We need to simulate the rotting process where:
- All initially rotten oranges start rotting their neighbors simultaneously
- Each minute, rotten oranges spread to adjacent fresh oranges
- We need to find the minimum time for all oranges to rot (or -1 if impossible)

The key insight is that we start BFS from ALL rotten oranges at once, not just one.
This simulates the parallel rotting process happening across the grid.
*/

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]int, rows)

	for i := range rows {
		visited[i] = make([]int, cols)
	}

	queue := [][]int{}

	// STEP 1: Find all initially rotten oranges and add them to queue
	// These are our starting points for BFS - they begin rotting at time 0
	for row := range rows {
		for col := range cols {
			if grid[row][col] == 2 {
				queue = append(queue, []int{
					row, col, 0, // [row, col, time_when_this_orange_rotted]
				})
				visited[row][col] = 2 // Mark as processed
			}
		}
	}

	ans := 0 // Track the maximum time needed

	// STEP 2: BFS to simulate the rotting process
	// Process oranges level by level (minute by minute)
	for len(queue) > 0 {
		// Dequeue the current rotten orange
		node := queue[0]
		queue = queue[1:]
		r, c, t := node[0], node[1], node[2]

		// Update the maximum time - this will be our final answer
		ans = max(ans, t)

		// Check all 4 adjacent cells manually

		// UP: Check the cell above (r-1, c)
		if r-1 >= 0 && grid[r-1][c] == 1 && visited[r-1][c] == 0 {
			visited[r-1][c] = 1
			grid[r-1][c] = 2
			queue = append(queue, []int{r - 1, c, t + 1})
		}

		// RIGHT: Check the cell to the right (r, c+1)
		if c+1 < cols && grid[r][c+1] == 1 && visited[r][c+1] == 0 {
			visited[r][c+1] = 1
			grid[r][c+1] = 2
			queue = append(queue, []int{r, c + 1, t + 1})
		}

		// DOWN: Check the cell below (r+1, c)
		if r+1 < rows && grid[r+1][c] == 1 && visited[r+1][c] == 0 {
			visited[r+1][c] = 1
			grid[r+1][c] = 2
			queue = append(queue, []int{r + 1, c, t + 1})
		}

		// LEFT: Check the cell to the left (r, c-1)
		if c-1 >= 0 && grid[r][c-1] == 1 && visited[r][c-1] == 0 {
			visited[r][c-1] = 1
			grid[r][c-1] = 2
			queue = append(queue, []int{r, c - 1, t + 1})
		}
	}

	// STEP 3: Check if any fresh oranges remain unrotted
	// If yes, return -1 (impossible to rot all oranges)
	for row := range rows {
		for col := range cols {
			if grid[row][col] == 1 {
				return -1 // Found a fresh orange that couldn't be reached
			}
		}
	}

	// Return the maximum time needed to rot all oranges
	return ans
}

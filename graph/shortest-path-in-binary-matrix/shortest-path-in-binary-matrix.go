package main

type pair struct {
	row      int
	col      int
	distance int
}

// 1091. Shortest Path in Binary Matrix
func shortestPathBinaryMatrix(grid [][]int) int {
	// If the starting cell is blocked, we can't start the path
	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)
	distances := make([][]int, n)

	// Initialize the distances matrix with a large value (infinity equivalent)
	for i := range n {
		distances[i] = make([]int, n)

		for j := range n {
			distances[i][j] = 1e9
		}
	}

	// set the starting source node's distance to 0
	distances[0][0] = 1

	// Initialize the BFS queue with the starting cell (0, 0) and distance 1
	queue := []pair{
		{
			row:      0,
			col:      0,
			distance: 1,
		},
	}

	// 8-directional movement: up, down, left, right, and diagonals
	delRows := []int{-1, 0, 1}
	delCols := []int{-1, 0, 1}

	// BFS traversal using a queue
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue the front element

		// Explore all 8 neighboring cells
		for _, row := range delRows {
			for _, col := range delCols {
				r := node.row + row
				c := node.col + col

				// Check if the new coordinates are within bounds
				if r >= 0 && r < n && c >= 0 && c < n {
					// Calculate the new distance if we move to this cell
					distance := node.distance + 1

					// If the cell is not blocked and we found a shorter path, update it
					if grid[r][c] == 0 && distance < distances[r][c] {
						distances[r][c] = distance
						// Enqueue the neighboring cell for further exploration
						queue = append(queue, pair{
							row:      r,
							col:      c,
							distance: distance,
						})
					}
				}
			}
		}
	}

	// If the destination cell (bottom-right) is still unreachable, return -1
	if distances[n-1][n-1] == 1e9 {
		return -1
	}

	// Return the shortest distance to the bottom-right corner
	return distances[n-1][n-1]
}

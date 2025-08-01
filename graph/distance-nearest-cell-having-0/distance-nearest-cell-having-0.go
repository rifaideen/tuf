package main

// updateMatrix returns a matrix where each cell contains the shortest distance
// from that cell to the nearest 0 in the original binary matrix.
//
// Intuition:
//   - This problem is a classic "multi-source shortest path" on a grid.
//   - Instead of starting from one point, we start from *all 0s* simultaneously.
//   - We use BFS (Breadth-First Search) because it naturally explores all nodes
//     level by level, which ensures that the first time we reach a cell,
//     we've found the shortest path to it.
//
// Why BFS?
//   - If we tried using DFS or checking each 1 individually, we'd risk taking
//     longer paths or doing redundant work.
//   - BFS from all 0s at once simulates a "wave" spreading out over the grid.
//   - Each step of the wave increases the distance by 1, which matches the problem's
//     definition of distance (Manhattan-style, adjacent moves only).
//
// Steps:
// 1. Initialize a visited matrix to keep track of which cells have been processed.
// 2. Initialize a distance matrix to store the result — the shortest distance to a 0.
// 3. Find all cells with value 0 — these are our starting points.
//   - Mark them as visited.
//   - Add them to the queue with distance 0 (since they are 0s themselves).
//
// 4. Begin BFS:
//   - For each cell dequeued, record its distance in the result matrix.
//   - Explore its 4 neighbors (up, down, left, right).
//   - If a neighbor hasn't been visited, mark it visited, and enqueue it
//     with distance = current_steps + 1.
//     5. As BFS progresses, the wave moves outward from all 0s, updating distances
//     in increasing order (0, then 1, then 2, etc.).
//
// This guarantees that each 1 gets the distance to the *closest* 0, because
// BFS always reaches a cell via the shortest path first.
//
// Example:
//
//	Input: [[0,1],[1,1]]
//	BFS starts at (0,0) → reaches (0,1) and (1,0) at step 1 →
//	then reaches (1,1) at step 2 → result: [[0,1],[1,2]]
//
// The algorithm efficiently computes the answer in O(m*n) time,
// since each cell is processed exactly once.
func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	visited := make([][]int, rows)
	distance := make([][]int, rows)

	for i := range visited {
		visited[i] = make([]int, cols)
		distance[i] = make([]int, cols)
	}

	queue := [][]int{}

	// Start BFS from all 0s simultaneously
	// All 0s are added to the queue with distance 0
	for row := range rows {
		for col := range cols {
			if mat[row][col] == 0 {
				visited[row][col] = 1
				queue = append(queue, []int{row, col, 0})
			}
		}
	}

	// Process the queue using BFS
	// Each time we visit a cell, we set its distance and propagate outward
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		row, col, steps := front[0], front[1], front[2]

		visited[row][col] = 1
		distance[row][col] = steps

		// Explore upward neighbor
		if row-1 >= 0 && visited[row-1][col] == 0 {
			visited[row-1][col] = 1
			queue = append(queue, []int{row - 1, col, steps + 1})
		}

		// Explore downward neighbor
		if row+1 < rows && visited[row+1][col] == 0 {
			visited[row+1][col] = 1
			queue = append(queue, []int{row + 1, col, steps + 1})
		}

		// Explore left neighbor
		if col-1 >= 0 && visited[row][col-1] == 0 {
			visited[row][col-1] = 1
			queue = append(queue, []int{row, col - 1, steps + 1})
		}

		// Explore right neighbor
		if col+1 < cols && visited[row][col+1] == 0 {
			visited[row][col+1] = 1
			queue = append(queue, []int{row, col + 1, steps + 1})
		}
	}

	return distance
}

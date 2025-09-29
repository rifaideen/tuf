package main

import (
	"math"
)

// maxChocolates returns the maximum chocolates two robots can collect
// starting from top-left (0,0) and top-right (0, n-1) of an n x n grid.
// Both move down simultaneously, one row at a time, and can adjust column by -1, 0, or +1.
func maxChocolates(grid [][]int) int {
	n := len(grid)
	// Start recursion from row 0, with robot1 at col 0 and robot2 at col n-1
	return f(0, 0, n-1, n, grid)
}

// f computes the max chocolates collectable from row `i` to the bottom,
// given robot1 is at column `j1` and robot2 at column `j2`.
func f(i, j1, j2, n int, grid [][]int) int {
	// If either robot goes out of bounds, return a very small number (invalid path)
	if j1 < 0 || j1 >= n || j2 < 0 || j2 >= n {
		return -1e9
	}

	// Base case: we're at the last row (i = n-1)
	if i == n-1 {
		if j1 == j2 {
			// Both robots are on the same cell → count chocolate only once
			return grid[i][j1]
		}
		// Different cells → collect both
		return grid[i][j1] + grid[i][j2]
	}

	// Each robot can move to 3 possible columns in the next row: -1, 0, +1
	// So total 3x3 = 9 combinations of moves
	delRows := []int{-1, 0, 1} // actually column deltas for robot1 (misnamed, but used as dj1)
	delCols := []int{-1, 0, 1} // actually column deltas for robot2 (misnamed, but used as dj2)

	maxValue := math.MinInt // track best total from this state

	// Try all 9 possible move combinations for the two robots
	for _, dr := range delRows {
		for _, dc := range delCols {
			ans := 0

			// Collect chocolates at current row before moving down
			if j1 == j2 {
				// Same cell → pick once
				ans = grid[i][j1] + f(i+1, j1+dr, j2+dc, n, grid)
			} else {
				// Different cells → pick both
				ans = grid[i][j1] + grid[i][j2] + f(i+1, j1+dr, j2+dc, n, grid)
			}

			// Update the best result among all move options
			maxValue = max(maxValue, ans)
		}
	}

	return maxValue
}

// maxChocolates returns the maximum chocolates two robots can collect
// starting from top-left (0,0) and top-right (0, n-1) of an n x n grid.
// Both move down simultaneously, one row at a time, and can adjust column by -1, 0, or +1.
func maxChocolatesMemoized(grid [][]int) int {
	n := len(grid)    // Number of rows
	m := len(grid[0]) // Number of columns

	memoized := make([][][]int, n)
	for i := range memoized {
		memoized[i] = make([][]int, m)
		for j := range memoized[i] {
			memoized[i][j] = make([]int, m)
			for k := range memoized[i][j] {
				memoized[i][j][k] = -1
			}
		}
	}

	// Start from row 0, Alice at col 0 and Bob at col m-1
	return fMemoized(0, 0, m-1, n, m, grid, &memoized)
}

func fMemoized(i, j1, j2, n, m int, grid [][]int, memoized *[][][]int) int {
	// If either robot is out of bounds
	if j1 < 0 || j1 >= m || j2 < 0 || j2 >= m {
		return -1e9
	}

	// Base case: last row
	if i == n-1 {
		if j1 == j2 {
			return grid[i][j1]
		}
		return grid[i][j1] + grid[i][j2]
	}

	if (*memoized)[i][j1][j2] != -1 {
		return (*memoized)[i][j1][j2]
	}

	delRows := []int{-1, 0, 1} // movement options
	delCols := []int{-1, 0, 1} // movement options

	maxValue := int(-1e5)

	for _, dr := range delRows {
		for _, dc := range delCols {
			ans := 0

			if j1 == j2 {
				ans = grid[i][j1] + fMemoized(i+1, j1+dr, j2+dc, n, m, grid, memoized)
			} else {
				ans = grid[i][j1] + grid[i][j2] + fMemoized(i+1, j1+dr, j2+dc, n, m, grid, memoized)
			}

			maxValue = max(maxValue, ans)
		}
	}

	(*memoized)[i][j1][j2] = maxValue
	return maxValue
}

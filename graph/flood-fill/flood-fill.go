package main

// floodFill performs a flood fill on a 2D image starting from the pixel at (sr, sc),
// changing all connected pixels of the same initial color to the new 'color'.
// This is similar to the "bucket fill" tool in paint programs.
//
// Parameters:
//   - image: 2D grid of integers representing pixel colors.
//   - sr: starting row index.
//   - sc: starting column index.
//   - color: the new color to fill with.
//
// Returns:
//   - The modified image after flood fill.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows := len(image)
	cols := len(image[0])

	// If the starting pixel already has the target color, no need to do anything.
	// This also prevents infinite recursion since we'd keep revisiting the same pixel.
	if image[sr][sc] == color {
		return image
	}

	// Store the original color of the starting pixel.
	// We'll use this to determine which pixels are part of the connected region.
	oldColor := image[sr][sc]

	// visited matrix to keep track of which pixels have been processed.
	// Prevents revisiting and redundant work (especially important to avoid cycles).
	visited := make([][]int, rows)
	for row := range rows {
		visited[row] = make([]int, cols)
	}

	// dfs is a nested recursive function that explores all 4-connected pixels
	// (up, down, left, right) that have the oldColor and haven't been visited yet.
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// Mark current pixel as visited and update its color.
		visited[row][col] = 1
		image[row][col] = color

		// Explore the four adjacent neighbors (up, down, left, right)

		// Move up
		if row-1 >= 0 && image[row-1][col] == oldColor && visited[row-1][col] == 0 {
			dfs(row-1, col)
		}

		// Move down
		if row+1 < rows && image[row+1][col] == oldColor && visited[row+1][col] == 0 {
			dfs(row+1, col)
		}

		// Move left
		if col-1 >= 0 && image[row][col-1] == oldColor && visited[row][col-1] == 0 {
			dfs(row, col-1)
		}

		// Move right
		if col+1 < cols && image[row][col+1] == oldColor && visited[row][col+1] == 0 {
			dfs(row, col+1)
		}
	}

	// Start the DFS from the given starting pixel (sr, sc)
	dfs(sr, sc)

	// Return the modified image after flood fill
	return image
}

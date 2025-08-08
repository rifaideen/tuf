package main

func numIslands(grid [][]string) int {
	ans := 0
	rows := len(grid)
	cols := len(grid[0])

	// Create visited array to keep track of visited nodes
	visited := make([][]int, rows)

	for row := range rows {
		visited[row] = make([]int, cols)
	}

	for row := range rows {
		for col := range cols {
			if grid[row][col] == "L" && visited[row][col] == 0 {
				ans++
				bfs(row, col, rows, cols, grid, &visited)
			}
		}
	}

	return ans
}

func bfs(r, c, rows, cols int, grid [][]string, visited *[][]int) {
	queue := [][]int{
		{r, c},
	}

	(*visited)[r][c] = 1

	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		(*visited)[row][col] = 1

		// move top
		if row-1 >= 0 {
			if (*visited)[row-1][col] == 0 && grid[row-1][col] == "L" {
				(*visited)[row-1][col] = 1
				queue = append(queue, []int{row - 1, col})
			}
		}

		// move top left
		if row-1 >= 0 && col-1 >= 0 {
			if (*visited)[row-1][col-1] == 0 && grid[row-1][col-1] == "L" {
				(*visited)[row-1][col-1] = 1
				queue = append(queue, []int{row - 1, col - 1})
			}
		}

		// move top right
		if row-1 >= 0 && col+1 < cols {
			if (*visited)[row-1][col+1] == 0 && grid[row-1][col+1] == "L" {
				(*visited)[row-1][col+1] = 1
				queue = append(queue, []int{row - 1, col + 1})
			}
		}

		// move bottom
		if row+1 < rows {
			if (*visited)[row+1][col] == 0 && grid[row+1][col] == "L" {
				(*visited)[row+1][col] = 1
				queue = append(queue, []int{row + 1, col})
			}
		}

		// move bottom left
		if row+1 < rows && col-1 >= 0 {
			if (*visited)[row+1][col-1] == 0 && grid[row+1][col-1] == "L" {
				(*visited)[row+1][col-1] = 1
				queue = append(queue, []int{row + 1, col - 1})
			}
		}

		// move bottom right
		if row+1 < rows && col+1 < cols {
			if (*visited)[row+1][col+1] == 0 && grid[row+1][col+1] == "L" {
				(*visited)[row+1][col+1] = 1
				queue = append(queue, []int{row + 1, col + 1})
			}
		}

		// move left
		if col-1 >= 0 {
			if (*visited)[row][col-1] == 0 && grid[row][col-1] == "L" {
				(*visited)[row][col-1] = 1
				queue = append(queue, []int{row, col - 1})
			}
		}

		// move right
		if col+1 < cols {
			if (*visited)[row][col+1] == 0 && grid[row][col+1] == "L" {
				(*visited)[row][col+1] = 1
				queue = append(queue, []int{row, col + 1})
			}
		}
	}

}

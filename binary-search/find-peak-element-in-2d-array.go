package main

func findPeakGrid(mat [][]int) []int {
	cols := len(mat[0])

	left := 0
	right := cols - 1

	for left <= right {
		mid := (left + right) / 2
		row := getRow(mat, mid)

		prev := -1
		next := -1

		if mid-1 >= 0 {
			prev = mat[row][mid-1]
		}

		if mid+1 < cols {
			next = mat[row][mid+1]
		}

		if mat[row][mid] > prev && mat[row][mid] > next {
			return []int{row, mid}
		}

		if mat[row][mid] < prev {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// dummy return statement to ful-fil the return type
	return []int{-1, -1}
}

func getRow(mat [][]int, mid int) int {
	maxValue := -1
	index := -1

	for row := range mat {
		if mat[row][mid] > maxValue {
			maxValue = mat[row][mid]
			index = row
		}
	}

	return index
}

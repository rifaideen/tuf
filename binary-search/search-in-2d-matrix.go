package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	for row := range rows {
		if matrix[row][0] <= target && target <= matrix[row][cols-1] {
			return searchTarget(matrix[row], target)
		}
	}

	return false
}

func searchTarget(row []int, target int) bool {
	left := 0
	right := len(row) - 1

	for left <= right {
		mid := (left + right) / 2

		if row[mid] == target {
			return true
		}

		if row[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

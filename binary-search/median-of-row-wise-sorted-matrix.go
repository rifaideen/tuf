package main

import "math"

func medianOfRowiseSortedMatrix(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	left := math.MaxInt
	right := math.MinInt

	for _, row := range matrix {
		left = min(left, row[0])
		right = max(right, row[cols-1])
	}

	required := (rows * cols) / 2

	for left <= right {
		mid := (left + right) / 2

		if countSmallEqual(matrix, rows, cols, mid) <= required {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func countSmallEqual(matrix [][]int, rows, cols, mid int) int {
	var count int

	for row := range rows {
		count += upperBound(matrix[row], cols, mid)
	}

	return count
}

func upperBound(matrix []int, cols, expected int) int {
	left := 0
	right := cols - 1
	ans := cols

	for left <= right {
		mid := (left + right) / 2

		if matrix[mid] > expected {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

package main

func MaxOnesRow(matrix [][]int) int {
	maxCount := 0
	index := -1

	for i := range len(matrix) {
		// using binary search to count the number of 1s
		count := countOnes(matrix[i])

		// updating the max count only when the count is greater then max count
		if count > maxCount {
			maxCount = count
			index = i
		}
	}

	return index
}

func countOnes(row []int) int {
	left := 0
	right := len(row) - 1

	for left <= right {
		mid := (left + right) / 2

		if row[mid] == 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return (len(row) - 1) - right
}

package main

func FindKthPositive(nums []int, k int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// avoiding interger overflow
		mid := left + (right-left)/2
		// calculate missing numbers
		missing := nums[mid] - (mid + 1) // mid + 1 is the expected value in the current position

		if missing < k { // we still have less number of missing numbers, move to right
			left = mid + 1
		} else { // we have more than required missing numbers, move to left
			right = mid - 1
		}
	}

	return left + k // left strictly less than k and adding left + k gives the result
}

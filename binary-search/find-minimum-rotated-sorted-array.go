package main

import "math"

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	ans := math.MaxInt

	for left <= right {
		mid := (left + right) / 2

		// identify which part is sorted
		if nums[left] <= nums[mid] {
			// left part is sorted, pickup the minimum value
			ans = min(ans, nums[left])

			// pickuped up the lowest value in the left side, move to right
			left = mid + 1
		} else {
			// right part is sorted, pickup the minimum value
			ans = min(ans, nums[mid])

			// picked up the lowest value in the right side, move to left
			right = mid - 1
		}
	}

	return ans
}

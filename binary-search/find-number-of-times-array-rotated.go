package main

import "math"

// find number of times the array is rotated
func FindRotatedCount(nums []int) int {
	left := 0
	right := len(nums) - 1
	ans := math.MaxInt
	index := -1

	for left <= right {
		mid := (left + right) / 2

		// search space is already sorted, nums[left] will be the smaller in search space
		if nums[left] <= nums[right] {
			if nums[left] < ans {
				// ans = nums[left]
				index = left
			}

			break
		}

		// identify which part is sorted. left or right?
		if nums[left] <= nums[mid] { // left side is sorted
			if nums[left] < ans {
				ans = nums[left]
				index = left
			}

			// picked up lowest index, move the pointer to the right
			left = mid + 1
		} else { // right side is sorted
			if nums[mid] < ans {
				ans = nums[mid]
				index = mid
			}

			// picked up lowest index, move the pointer to the left
			right = mid - 1
		}
	}

	return index
}

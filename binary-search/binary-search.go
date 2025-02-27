package main

// iterative approach
func BinarySearch(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func RecursiveBinarySearch(nums []int, target int) int {
	return recursiveBS(nums, 0, len(nums)-1, target)
}

func recursiveBS(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return recursiveBS(nums, left, right, target)
}

package main

func FindPeakElement(nums []int) int {
	n := len(nums)

	// when size is 1, he is the one.
	if n == 1 {
		return 0
	}

	// check if peak is in front
	if nums[0] > nums[1] {
		return 0
	}

	// check if peak is in back
	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	// start with shriked pointers to avoid edge cases
	left := 1
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		// check if mid is peak
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		// check if mid is in increasing order
		if nums[mid] < nums[mid+1] {
			// mid is in increasing order, move the pointer to the right
			left = mid + 1
		} else {
			// mid is in the decreasing order, move the pointer to the left
			right = mid - 1
		}
	}

	return -1
}

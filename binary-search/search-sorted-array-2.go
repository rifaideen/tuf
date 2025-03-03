package main

func SearchSortedArrayII(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target { // when current number is target, return the current index
			return mid
		}

		// shrink search space when left == mid == right and continue
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--

			continue
		}

		// identify which side is sorted, left or right?
		if nums[left] <= nums[mid] {
			// left side is sorted

			// check if target is within the range (within left side)
			if nums[left] <= target && target <= nums[mid] {
				// target is within the range (within left side), discard right side
				right = mid - 1
			} else {
				// target is outside the range (not in left side), discard left side
				left = mid + 1
			}
		} else {
			// right side is sorted

			// check if target is within the range (within right side)
			if nums[mid] <= target && target <= nums[right] {
				// target is within the range (within right side), discard left side
				left = mid + 1
			} else {
				// target is outside the range (not in right side), discard right side
				right = mid - 1
			}
		}
	}

	// we could not find the target at all
	return -1
}

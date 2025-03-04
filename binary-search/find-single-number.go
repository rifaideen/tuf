package main

// find single element in the given array
// each elements appears twice, except one number
func FindSingleNumber(nums []int) int {
	n := len(nums)

	// when size is 1, it is the single number
	if n == 1 {
		return nums[0]
	}

	// when first and second not same
	if nums[0] == nums[1] {
		return nums[0]
	}

	// when last and element before last not same
	if nums[n-1] == nums[n-2] {
		return nums[n-1]
	}

	// start the pointers with one step narrowed to avoid index range issue
	var left = 1
	var right = n - 2

	for left <= right {
		var mid = (left + right) / 2

		// if the value before and after the mid is not same, he is the one. catch him
		if nums[mid-1] != nums[mid] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if mid%2 == 1 && nums[mid] == nums[mid-1] { // we're on left
			left = mid + 1
		} else if mid%2 == 0 && nums[mid] == nums[mid+1] { // we're on left
			left = mid + 1
		} else { // we're on right
			right = mid - 1
		}
	}

	return -1
}

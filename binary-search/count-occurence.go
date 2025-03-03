package main

func CountOccurence(nums []int, target int) int {
	left := first(nums, target)
	right := last(nums, target)

	return right - left + 1
}

func first(nums []int, target int) int {
	ans := -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func last(nums []int, target int) int {
	ans := -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

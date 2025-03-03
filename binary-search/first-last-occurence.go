package main

func FirstAndLastOccurence(nums []int, target int) []int {
	n := len(nums)
	first := findFirst(nums, target, 0, n-1)
	last := findLast(nums, target, 0, n-1)

	return []int{first, last}
}

func findFirst(nums []int, target, left, right int) int {
	first := -1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			first = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return first
}

func findLast(nums []int, target, left, right int) int {
	last := -1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return last
}

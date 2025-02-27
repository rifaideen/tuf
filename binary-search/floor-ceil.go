package main

// return floor and ceil of the given value in sorted nums array
//
// use the below formula when finding the floor and ceil on sorted numbers
// floor = nums[i] <= n
// ceil = nums[i] >= n
func FloorCeil(nums []int, x int) []int {
	n := len(nums)
	low := 0
	high := n - 1
	floor := -1
	ceil := -1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] >= x {
			ceil = nums[mid]
			high = mid - 1
		}

		if nums[mid] <= x {
			floor = nums[mid]
			low = mid + 1
		}
	}

	return []int{floor, ceil}
}

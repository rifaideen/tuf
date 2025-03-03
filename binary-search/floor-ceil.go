package main

// return floor and ceil of the given value in sorted nums array
//
// use the below formula when finding the floor and ceil on sorted numbers
// floor = nums[i] <= n
// ceil = nums[i] >= n
func FloorCeil(nums []int, x int) []int {
	return []int{
		floor(nums, x),
		ceil(nums, x),
	}
}

// floor = nums[i] <= n
func floor(nums []int, x int) int {
	n := len(nums)
	left := 0
	right := n - 1
	ans := -1

	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] <= x {
			ans = mid
			left = mid + 1 // answer might me too low, search more on the right side
		} else {
			right = mid - 1
		}
	}

	return ans
}

func ceil(nums []int, x int) int {
	n := len(nums)
	left := 0
	right := n - 1
	ans := n

	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] >= x {
			ans = mid
			right = mid - 1 // answer might be too high, search on the left side
		} else {
			left = mid + 1
		}
	}

	return ans
}

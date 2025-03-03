package main

func SearchInsert(nums []int, target int) int {
	ans := len(nums)
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

package main

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithKDistinctH(nums, k) - subarraysWithKDistinctH(nums, k-1)
}

func subarraysWithKDistinctH(nums []int, k int) int {
	if k <= 0 {
		return 0
	}

	n := len(nums)
	left := 0
	right := 0
	count := 0
	frequency := map[int]int{}

	for right < n {
		frequency[nums[right]]++

		for len(frequency) > k {
			frequency[nums[left]] -= 1

			if frequency[nums[left]] == 0 {
				delete(frequency, nums[left])
			}

			left++
		}

		// if len(frequency) <= k {
		count += right - left + 1
		// }

		right++
	}

	return count
}

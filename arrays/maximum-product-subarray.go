package main

import "math"

func MaxProductSubarray(nums []int) int {
	ans := -math.MaxInt
	prefix := 1
	suffix := 1
	n := len(nums)

	for i := 0; i < n; i++ {
		// reset prefix when it becomes zero
		if prefix == 0 {
			prefix = 1
		}

		// reset suffix when it becomes zero
		if suffix == 0 {
			suffix = 1
		}

		prefix *= nums[i]
		suffix *= nums[(n-1)-i]

		// choose the maximum between answer, prefix and suffix
		ans = max(ans, max(prefix, suffix))
	}

	return ans
}

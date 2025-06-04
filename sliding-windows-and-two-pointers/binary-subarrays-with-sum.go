package main

func numSubarraysWithSum(nums []int, goal int) int {
	return numSubarraysWithSumHelper(nums, goal) - numSubarraysWithSumHelper(nums, goal-1)
}

func numSubarraysWithSumHelper(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}

	left := 0
	right := 0
	ans := 0
	sum := 0

	for right < len(nums) {
		sum += nums[right]

		for sum > goal {
			sum -= nums[left]
			left++
		}

		ans += right - left + 1
		right++
	}

	return ans
}

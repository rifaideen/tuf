package main

func numberOfSubarrays(nums []int, k int) int {
	return numberOfSubarraysHelper(nums, k) - numberOfSubarraysHelper(nums, k-1)
}

func numberOfSubarraysHelper(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	left := 0
	right := 0
	sum := 0
	ans := 0

	for right < len(nums) {
		sum += nums[right] % 2

		for sum > k {
			sum -= nums[left] % 2
			left++
		}

		ans += right - left + 1

		right++
	}

	return ans
}

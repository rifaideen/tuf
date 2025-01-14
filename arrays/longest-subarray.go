package main

func longestSubarray(nums []int, k int) int {
	left := 0
	curr := 0
	ans := 0

	for right := 0; right < len(nums); right++ {
		curr += nums[right]

		for curr > k {
			curr -= nums[left]
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

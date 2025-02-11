package main

func LongestSubArraySumZero(nums []int) int {
	sum := 0
	maxLen := 0
	prefixSum := map[int]int{}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum == 0 {
			maxLen = i + 1
		} else {
			if left, ok := prefixSum[sum]; ok {
				maxLen = max(i-left, maxLen)
			} else {
				prefixSum[sum] = i
			}
		}

	}

	return maxLen
}

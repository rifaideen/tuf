package main

func LongestSubArraySumZero(nums []int) int {
	sum := 0
	maxLen := 0
	prefixSum := map[int]int{} // key: sum, value: index

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum == 0 { // when sum is zero, set maxlen to i + 1 (since i is 0 based we have to add 1)
			maxLen = i + 1
		} else {
			// when sum already exit in map, calculate max len, else add it to map
			if left, ok := prefixSum[sum]; ok {
				maxLen = max(i-left, maxLen)
			} else {
				prefixSum[sum] = i
			}
		}

	}

	return maxLen
}

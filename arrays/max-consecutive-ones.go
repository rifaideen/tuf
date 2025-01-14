package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	curr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curr++

			if curr > max {
				max = curr
			}
		} else {
			curr = 0
		}
	}

	return max
}

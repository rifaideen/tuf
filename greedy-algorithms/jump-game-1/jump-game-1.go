package main

func canJump(nums []int) bool {
	maxIndex := 0

	// During each iteration, I will be keep track of what is the maximum index I can reach
	for index, num := range nums {
		// I cannot cross the max index, stop here. I can cross upto max-index
		if index > maxIndex {
			return false
		}

		// keep the max index upto date.
		maxIndex = max(maxIndex, index+num)
	}

	return true
}

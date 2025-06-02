package main

func longestOnes3(nums []int, k int) int {
	n := len(nums)
	left := 0
	maxLen := 0
	zeroes := 0

	// with two-pointers left and right starts with 0th index
	for right := 0; right < n; right++ {
		// when num is 0, increment the zeroes counter
		if nums[right] == 0 {
			zeroes++
		}

		// the moment zeroes goes beyond k, move the left pointer and when left points to 0, decrement the zeroes count
		if zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}

			left++
		}

		// when zeroes less than or equals to k, update the maxLen
		if zeroes <= k {
			maxLen = max(maxLen, right-left+1)
		}
	}

	return maxLen
}

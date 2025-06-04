package main

func characterReplacement(s string, k int) int {
	// left and right pointer
	left := 0
	right := 0

	// frequency array to store the frequency count for each alphabet in s
	frequency := [26]int{}
	// we need to keep track of max frequency to make sure if the window is valid
	maxFrequency := 0

	// to keep track of max len of valid substring we have seen so far
	maxLen := 0

	for right < len(s) {
		// incremement the frequency
		frequency[s[right]-'A']++
		// calculate the max frequency
		maxFrequency = max(maxFrequency, frequency[s[right]-'A'])

		// shrink the window when the window crosses k
		if (right-left+1)-maxFrequency > k {
			// reduce the current left alphabet's frequency
			frequency[s[left]-'A']--
			// move the left pointer
			left++
		}

		// check if window and the max frequency is within k and calculate the answer
		if (right-left+1)-maxFrequency <= k {
			maxLen = max(maxLen, right-left+1)
		}

		right++
	}

	return maxLen
}

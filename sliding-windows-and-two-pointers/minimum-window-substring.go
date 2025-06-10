package main

import "math"

func minWindow(s string, t string) string {
	// two-pointers in sliding window
	left := 0
	right := 0

	// keep track of min length start index
	minIndex := math.MaxInt
	// keep track of how many items we have found among t items in frequency
	count := 0
	// start index with -1, when start index is still at -1, we haven't seen any t in s
	startIndex := -1
	// store the unicode frequency (a-z and A-Z) count
	frequency := [256]int{}

	// pre-calculate frequency with t
	for i := 0; i < len(t); i++ {
		frequency[t[i]]++
	}

	// start the sliding-window
	for right < len(s) {
		// if the current char in frequency is > 0, we got 1 item from t, so increment the count by 1
		if frequency[s[right]] > 0 {
			count++
		}

		// regardless of we see it or not, we should decrease current character from the frequency
		frequency[s[right]]--

		// while count equals to size of t, shrink the window by moving left pointer
		for count == len(t) {
			// check if current window size is less than existing min-index
			// update the min-index and set the start index to left
			if (right - left + 1) < minIndex {
				minIndex = right - left + 1
				startIndex = left
			}

			// as we move the left, we should put it back to the frequency
			frequency[s[left]]++

			// at any point when the frequency of left > 0, decrease the count by 1
			if frequency[s[left]] > 0 {
				count--
			}

			// move the left pointer
			left++
		}

		right++
	}

	// did we ever updated the starting index? if not, return empty string
	if startIndex == -1 {
		return ""
	}

	// send the substring from the start index to min index
	return s[startIndex : startIndex+minIndex]
}

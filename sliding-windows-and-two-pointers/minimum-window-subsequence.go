package main

import (
	"math"
)

// minWindowSubsequence finds the minimum window in s1 such that it contains all characters of s2 in a subsequence manner.
func minWindowSubsequence(s1, s2 string) string {
	j := 0                   // points to current position in s2
	minLength := math.MaxInt // initialize minLength with maximum possible integer value
	startIndex := -1         // start index of the minimum window

	// iterate over s1 using variable i as the pointer
	for i := 0; i < len(s1); i++ {
		// if current character in s1 matches the current character in s2, increment j
		if s1[i] == s2[j] {
			j++

			// if all characters of s2 have been matched (j equals length of s2),
			if j == len(s2) {
				end := i + 1 // set end as the position after the current character in s1
				// move the j pointer back to find the minimal length window that contains subsequence s2
				j--

				// iterate backwards over s1 using variable i and s2 using variable j
				for j >= 0 {
					if s1[i] == s2[j] { // if characters match, decrement j
						j--
					}

					i-- // move to the previous character in s1
				}

				i++ // adjust i to point to the start of the found window, this is to reverse the last decrement inside the j >= 0 loop
				j++

				// update minLength and startIndex if a smaller valid window is found
				if end-i < minLength {
					minLength = end - i
					startIndex = i
				}
			}
		}
	}

	// return an empty string if no valid window was found, otherwise return the minimum window substring
	if startIndex == -1 {
		return ""
	}

	return s1[startIndex : startIndex+minLength]
}

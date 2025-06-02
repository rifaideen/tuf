package main

func characterReplacement(s string, k int) int {
	left := 0
	replaced := 0
	maxLen := 0

	for right := range len(s) {
		if s[left] != s[right] {
			if replaced <= k {
				replaced++
			} else {
				break
			}
		}

		if replaced <= k {
			maxLen = max(maxLen, right-left+1)
		}
	}

	return maxLen
}

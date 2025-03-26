package main

import (
	"math"
	"strings"
)

// using horizontal scanning
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// take the first item as prefix
	prefix := strs[0]

	// compare prefix with rest of the strs array
	for i := 1; i < len(strs); i++ {
		// if the string does not have the prefix, we cut one character from the right and compare again
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func longestCommonPrefixBS(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 1 to min length of strs
	left := 1
	right := math.MaxInt

	for _, str := range strs {
		right = min(right, len(str))
	}

	for left <= right {
		mid := (left + right) / 2

		if isCommonPrefix(strs, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return strs[0][:right]
}

func isCommonPrefix(strs []string, mid int) bool {
	prefix := strs[0][:mid]

	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			return false
		}
	}

	return true
}

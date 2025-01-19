package main

import "github.com/emirpasic/gods/sets/hashset"

func longestConsecutive(nums []int) int {
	longest := 1
	s := hashset.New()

	for i := 0; i < len(nums); i++ {
		s.Add(nums[i])
	}

	for _, num := range s.Values() {
		if !s.Contains(num.(int) - 1) {
			count := 1
			start := num.(int)

			for s.Contains(start + 1) {
				start += 1
				count += 1
			}

			longest = max(longest, count)
		}
	}

	return longest
}

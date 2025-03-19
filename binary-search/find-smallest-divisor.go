package main

import (
	"math"
	"slices"
)

func SmallestDivisor(nums []int, threshold int) int {
	// start the pointer from 1 to Max(nums)
	left := 1
	right := slices.Max(nums)

	for left <= right {
		mid := (left + right) / 2

		// get the division count and compare with threshold
		if count(nums, float64(mid)) <= threshold {
			// within threshold, move the pointer to the left
			right = mid - 1
		} else {
			// outside the threshold, move the pointer to the right
			left = mid + 1
		}
	}

	return left
}

func count(nums []int, divisor float64) int {
	count := 0.0

	for _, num := range nums {
		count += math.Ceil(float64(num) / divisor)
	}

	return int(count)
}

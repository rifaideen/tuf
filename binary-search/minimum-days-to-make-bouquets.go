package main

import "slices"

// bloomDays - array of days to make a rose bloom
// m - the number of bouquets to make
// k - the number of adjacent flowers to use
func MinDays(bloomDays []int, m, k int) int {
	n := len(bloomDays)

	if n < m*k {
		return -1
	}

	left := slices.Min(bloomDays)
	right := slices.Max(bloomDays)

	for left <= right {
		mid := (left + right) / 2

		if possible(bloomDays, mid, m, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func possible(bloomDays []int, day, m, k int) bool {
	roses := 0
	bouquets := 0

	for _, days := range bloomDays {
		// we have rose bloomed within specic day value
		if days <= day {
			roses++
		} else { // we don't have continuous blooming
			// calculate number of bouquest we can make so far with the roses we have
			bouquets += roses / k
			// reset the counter
			roses = 0
		}
	}

	// calculate the number of bouquest with the remaining roses
	bouquets += roses / k

	// check if we made enough bouquets?
	return bouquets >= m
}

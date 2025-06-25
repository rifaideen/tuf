package main

import "slices"

// We solve this problem using two pointers, sorting the departures and arrivals
//
// the arrivals and departures are handled in ascending order
// if train comes first, we increment the platform and move left pointer, when it's time for departure of any train, we decrement the platform and move the right pointer
// the max platform we required is the minimal number of platfomrs required
func minPlatform(arrivals, departures []int) int {
	slices.Sort(arrivals)
	slices.Sort(departures)

	n := len(arrivals)
	left := 0
	right := 0
	platforms := 0
	minPlatforms := 0

	for left < n && right < n {
		if arrivals[left] <= departures[right] {
			platforms++
			left++
		} else {
			platforms--
			right++
		}

		minPlatforms = max(minPlatforms, platforms)
	}

	return minPlatforms
}

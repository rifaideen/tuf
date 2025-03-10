package main

import "slices"

func AggressiveCows(stalls []int, k int) int {
	// sorting to make sure we are able to calculate the distance without hassle
	slices.Sort(stalls)

	// starting from 1 to (max - min)
	left := 1
	right := stalls[len(stalls)-1] - stalls[0]

	for left <= right {
		mid := (left + right) / 2

		if checkCows(stalls, mid, k) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func checkCows(stalls []int, mid, k int) bool {
	// since we can place at least 1 cow in stall[0], we start with 1 and stalls[0]
	cows := 1
	last := stalls[0]

	for i := 1; i < len(stalls); i++ {
		// compare if we have enough distance to place the cow
		if (stalls[i] - last) >= mid {
			// cow placed and update the last location
			cows++
			last = stalls[i]
		}
	}

	// check if we have placed enough cows
	return cows >= k
}

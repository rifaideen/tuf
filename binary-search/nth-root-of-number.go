package main

import "math"

func NthRoot(n, m int) int {

	left := 1
	right := m

	for left <= right {
		mid := (left + right) / 2

		val := int(math.Pow(float64(mid), float64(n)))

		if val == m {
			return mid
		}

		if val > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

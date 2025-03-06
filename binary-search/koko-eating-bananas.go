package main

import (
	"math"
	"slices"
)

func MinEatingSpeed(piles []int, h int) int {
	left := 1
	right := slices.Max(piles)

	for left <= right {
		mid := (left + right) / 2
		total := countHours(piles, mid)

		if total <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func countHours(piles []int, hourly int) int {
	total := 0

	for banana := range piles {
		total += int(math.Ceil(float64(piles[banana]) / float64(hourly)))
	}

	return total
}

package main

import (
	"math"
)

func beautySum(s string) int {
	// running sum
	ans := 0

	// generate all the possible substring
	for i := 0; i < len(s); i++ {
		// store frequency map in incremental order
		frequency := map[byte]int{}

		for j := i; j < len(s); j++ {
			// calculate frequency
			frequency[s[j]]++

			most := 0
			least := math.MaxInt
			first := true

			// calculate the most and least for the overall frequency
			for _, v := range frequency {
				most = max(most, v)

				if first {
					least = v
					first = false
					continue
				}

				least = min(least, v)
			}

			// apply the beauty of substring formula to the running sum
			ans += most - least
		}
	}

	return ans
}

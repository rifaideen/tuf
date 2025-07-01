package main

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	ans = append(ans, intervals[0])

	for i := 1; i < len(intervals); i++ {
		x, y := intervals[i][0], intervals[i][1]

		stack := ans[len(ans)-1]

		if x <= stack[1] {
			stack[1] = max(stack[1], y)
		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

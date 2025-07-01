package main

import (
	"cmp"
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	// sort the intervals by end time
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	n := len(intervals)

	// we can perform at-least 1 meeting, hence start with 1
	meetings := 1
	// store the first interval's end time
	lastEndTime := intervals[0][1]

	for i := 1; i < n; i++ {
		// check if current interval's start time is greater than or equal to last interval's end time
		if intervals[i][0] >= lastEndTime {
			// update the meetings and set the last end time
			meetings++
			lastEndTime = intervals[i][1]
		}
	}

	// we know how many valid meetings we can perform, subtracting it from n give the min number of intervals to remove
	// solved using the same one room, n meetings problem
	return n - meetings
}

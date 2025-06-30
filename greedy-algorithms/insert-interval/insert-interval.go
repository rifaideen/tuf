package main

// We split the intervals into 3 parts:
// left side, which doesn't require merging, add it to answer
// center, which overlaps, which might require merging, expand the new interval
// right side, which doesn't require merging, add it to the answer
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	ans := [][]int{}
	i := 0

	// find the intervals which doesn't require expanding with new intervals
	// check if current interval ends before the new interval starts and push it to answer
	for i < n && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}

	// merge the overlapping intervals
	// check if current interval start time is within new interval's end time
	for i < n && intervals[i][0] <= newInterval[1] {
		// expand the new interval and put the minimum of current and new interval as the start time
		newInterval[0] = min(newInterval[0], intervals[i][0])
		// expand the new interval and put the maximum of current and new interval as the end time
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	// append the merged new interval
	ans = append(ans, newInterval)

	// append the remaining intervals which are left untouched after merging
	for i < n {
		ans = append(ans, intervals[i])
		i++
	}

	return ans
}

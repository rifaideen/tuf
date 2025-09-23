package main

// ninjaTraining: Entry point. Start from last day (n-1) with no previous task (-1).
// We want max points over n days, avoiding same task on consecutive days.
func ninjaTraining(n int, points [][]int) int {
	return f(n-1, -1, points)
}

// f(day, last, points): Recursive brute-force solution.
// Returns max points from day 0 to 'day', where 'last' was the task done on day-1.
// On day 0, we pick best task except 'last'.
// For other days, try all valid tasks today and recurse for yesterday.
func f(day, last int, points [][]int) int {
	if day == 0 { // Base case: first day
		maxValue := 0

		for i := range 3 { // Try each task on day 0
			if last != i { // Can't do the same as yesterday
				maxValue = max(maxValue, points[0][i])
			}
		}
		return maxValue
	}

	maxValue := 0
	for i := range 3 { // Try each task today
		if last != i { // Skip if same as yesterday
			value := points[day][i] + f(day-1, i, points) // Today's points + best from yesterday (with today's task as last)
			maxValue = max(maxValue, value)
		}
	}

	return maxValue
}

// ninjaTrainingMemoized: Same as above but uses memoization to avoid recomputation.
// memoized[day][last] stores best result up to 'day' when last task was 'last'.
// last can be 0,1,2 or 3 (3 means "no previous task", like start).
func ninjaTrainingMemoized(n int, points [][]int) int {
	memoized := make([][]int, n)

	for day := range memoized {
		memoized[day] = []int{-1, -1, -1, -1} // Initialize all states as uncomputed
	}

	return fMemoized(n-1, 3, points, &memoized) // Start from last day, no previous task (3 = none)
}

// fMemoized: Top-down DP with caching.
// If already computed, return cached result.
// Otherwise, compute recursively and store.
func fMemoized(day, last int, points [][]int, memoized *[][]int) int {
	if day == 0 { // First day: pick best task except 'last'
		maxValue := 0
		for i := range 3 {
			if last != i {
				maxValue = max(maxValue, points[0][i])
			}
		}

		(*memoized)[day][last] = maxValue // Cache result

		return maxValue
	}

	// If already computed, return it
	if (*memoized)[day][last] != -1 {
		return (*memoized)[day][last]
	}

	maxValue := 0
	for i := range 3 {
		if last != i { // Can't repeat task
			value := points[day][i] + fMemoized(day-1, i, points, memoized)
			maxValue = max(maxValue, value)
		}
	}

	(*memoized)[day][last] = maxValue // Store before returning
	return maxValue
}

// ninjaTrainingTabulation: Bottom-up DP. No recursion.
// Fill table iteratively from day 0 to day n-1.
// tabulation[day][last] = max points up to 'day' if task 'last' was done on day-1.
// last=3 means no prior task (used for final answer).
func ninjaTrainingTabulation(n int, points [][]int) int {
	tabulation := make([][]int, n)

	for day := range tabulation {
		tabulation[day] = []int{-1, -1, -1, -1}
	}

	// Base case: Day 0
	// If previous task was 0 → can pick max of task 1 or 2
	tabulation[0][0] = max(points[0][1], points[0][2])
	// If previous was 1 → pick max of 0 or 2
	tabulation[0][1] = max(points[0][0], points[0][2])
	// If previous was 2 → pick max of 0 or 1
	tabulation[0][2] = max(points[0][0], points[0][1])
	// If no previous task (like start), pick best of all three
	tabulation[0][3] = max(points[0][0], points[0][1], points[0][2])

	// Fill the table for remaining days
	for day := 1; day < n; day++ {
		for last := range 4 { // last = 0,1,2,3 (3 = no task)
			for task := range 3 { // try doing each task today
				if last != task { // can't repeat task
					// Update: points from doing 'task' today + best from yesterday ending with 'task'
					tabulation[day][last] = max(tabulation[day][last], points[day][task]+tabulation[day-1][task])
				}
			}
		}
	}

	// Final answer: best total points ending at day n-1 with any starting condition (last=3 means no constraint)
	return tabulation[n-1][3]
}

// ninjaTrainingOptimized: Space-optimized version.
// Only keep track of previous day’s results.
// Avoids full 2D table → O(1) space instead of O(n).
func ninjaTrainingOptimized(n int, points [][]int) int {
	// previous[last] = best total points up to yesterday, if yesterday’s task was 'last'
	previous := make([]int, 4)

	// Initialize for day 0 (same as tabulation)
	previous[0] = max(points[0][1], points[0][2])
	previous[1] = max(points[0][0], points[0][2])
	previous[2] = max(points[0][0], points[0][1])
	previous[3] = max(points[0][0], points[0][1], points[0][2])

	// Process each day, updating only current state
	for day := 1; day < n; day++ {
		temp := make([]int, 4) // hold today’s results

		for last := range 4 { // what was yesterday’s task?
			for task := range 3 { // what task am I doing today?
				if last != task { // can't repeat
					// Today’s best ending with 'last': try doing 'task' today + yesterday’s result ending with 'task'
					temp[last] = max(temp[last], points[day][task]+previous[task])
				}
			}
		}

		previous = temp // move to next day
	}

	// Return best possible score after last day (with no restriction on yesterday’s task → index 3)
	return previous[3]
}

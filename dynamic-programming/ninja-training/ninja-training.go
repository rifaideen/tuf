package main

func ninjaTraining(n int, points [][]int) int {
	return f(n-1, -1, points)
}

func f(day, last int, points [][]int) int {
	if day == 0 {
		maxValue := 0

		for i := range 3 {
			if last != i {
				maxValue = max(maxValue, points[0][i])
			}
		}

		return maxValue
	}

	maxValue := 0

	for i := range 3 {
		if last != i {
			value := points[day][i] + f(day-1, i, points)

			maxValue = max(maxValue, value)
		}
	}

	return maxValue
}

func ninjaTrainingMemoized(n int, points [][]int) int {
	memoized := make([][]int, n)

	for day := range memoized {
		memoized[day] = []int{-1, -1, -1, -1}
	}

	return fMemoized(n-1, 3, points, &memoized)
}

func fMemoized(day, last int, points [][]int, memoized *[][]int) int {
	if day == 0 {
		maxValue := 0

		for i := range 3 {
			if last != i {
				maxValue = max(maxValue, points[0][i])
			}
		}

		(*memoized)[day][last] = maxValue

		return maxValue
	}

	if (*memoized)[day][last] != -1 {
		return (*memoized)[day][last]
	}

	maxValue := 0

	for i := range 3 {
		if last != i {
			value := points[day][i] + fMemoized(day-1, i, points, memoized)

			maxValue = max(maxValue, value)
		}
	}

	(*memoized)[day][last] = maxValue

	return maxValue
}

func ninjaTrainingTabulation(n int, points [][]int) int {
	tabulation := make([][]int, n)

	for day := range tabulation {
		tabulation[day] = []int{-1, -1, -1, -1}
	}

	tabulation[0][0] = max(points[0][1], points[0][2])
	tabulation[0][1] = max(points[0][0], points[0][2])
	tabulation[0][2] = max(points[0][0], points[0][1])
	tabulation[0][3] = max(points[0][0], points[0][1], points[0][2])

	for day := 1; day < n; day++ {
		for last := range 4 {
			maxValue := 0

			for task := range 3 {
				if last != task {
					maxValue = max(maxValue, points[day][task]+tabulation[day-1][task])
				}
			}

			tabulation[day][last] = maxValue
		}
	}

	return tabulation[n-1][3]
}

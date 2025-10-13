package main

func unboundedKnapsack(weights []int, val []int, n int, W int) int {
	return f(n-1, W, weights, val)
}

func f(index, target int, weights, val []int) int {
	if index == 0 {
		return (target / weights[0]) * val[0]
	}

	pick := 0

	if weights[index] <= target {
		pick = val[index] + f(index, target-weights[index], weights, val)
	}

	notPick := 0 + f(index-1, target, weights, val)

	return max(pick, notPick)
}

func unboundedKnapsackMemoized(weights []int, val []int, n int, W int) int {
	memoized := make([][]int, n)

	for i := range memoized {
		memoized[i] = make([]int, W+1)

		for j := range memoized[i] {
			memoized[i][j] = -1
		}
	}

	return fMemoized(n-1, W, weights, val, &memoized)
}

func fMemoized(index, target int, weights, val []int, memoized *[][]int) int {
	if index == 0 {
		return (target / weights[0]) * val[0]
	}

	if (*memoized)[index][target] != -1 {
		return (*memoized)[index][target]
	}

	pick := 0

	if weights[index] <= target {
		pick = val[index] + fMemoized(index, target-weights[index], weights, val, memoized)
	}

	notPick := 0 + fMemoized(index-1, target, weights, val, memoized)

	(*memoized)[index][target] = max(pick, notPick)

	return (*memoized)[index][target]
}

func unboundedKnapsackTabulation(weights []int, val []int, n int, W int) int {
	tabulation := make([][]int, n)

	for i := range tabulation {
		tabulation[i] = make([]int, W+1)
	}

	for i := range W + 1 {
		tabulation[0][i] = (i / weights[0]) * val[0]
	}

	for i := 1; i < n; i++ {
		for target := range W + 1 {
			pick := 0

			if weights[i] <= target {
				pick = val[i] + tabulation[i][target-weights[i]]
			}

			notPick := tabulation[i-1][target]

			tabulation[i][target] = max(pick, notPick)
		}
	}

	return tabulation[n-1][W]
}

func unboundedKnapsackOptimized(weights []int, val []int, n int, W int) int {
	previous := make([]int, W+1)

	for target := range W + 1 {
		previous[target] = (target / weights[0]) * val[0]
	}

	for i := 1; i < n; i++ {
		current := make([]int, W+1)

		for target := range W + 1 {
			pick := 0

			if weights[i] <= target {
				pick = val[i] + current[target-weights[i]]
			}

			notPick := previous[target]

			current[target] = max(pick, notPick)
		}

		previous = append([]int{}, current...)
	}

	return previous[W]
}

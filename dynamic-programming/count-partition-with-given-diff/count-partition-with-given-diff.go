package main

func countPartitions(n int, diff int, arr []int) int {
	totalSum := 0

	for i := range arr {
		totalSum += arr[i]
	}

	if (totalSum-diff) < 0 || (totalSum-diff)%2 == 1 {
		return 0
	}

	// // eq1
	// s1 + s2 = totalSum

	// // eq2
	// s1 - s2 = diff
	// s1 = diff + s2

	// diff + s2 + s2 = totalSum
	// diff + 2s2 = totalSum
	// 2s2 = totalSum - diff
	// s2 = (totalSum - diff) / 2

	s2 := (totalSum - diff) / 2

	memoized := make([][]int, n)

	for i := range memoized {
		memoized[i] = make([]int, s2+1)

		for j := range memoized[i] {
			memoized[i][j] = -1
		}
	}

	return f(n-1, s2, arr, &memoized)
}

func f(index, target int, arr []int, memoized *[][]int) int {
	if index == 0 {
		// check if first item equals to target
		if arr[0] == 0 && target == 0 {
			return 2
		}

		if target == 0 || arr[0] == target {
			return 1
		}

		return 0
	}

	if (*memoized)[index][target] != -1 {
		return (*memoized)[index][target]
	}

	pick := 0

	if arr[index] <= target {
		pick = f(index-1, target-arr[index], arr, memoized)
	}

	notPick := f(index-1, target, arr, memoized)

	(*memoized)[index][target] = pick + notPick

	return (*memoized)[index][target]
}

func countPartitionsT(n int, diff int, arr []int) int {
	totalSum := 0

	for i := range arr {
		totalSum += arr[i]
	}

	if (totalSum-diff) < 0 || (totalSum-diff)%2 == 1 {
		return 0
	}

	// // eq1
	// s1 + s2 = totalSum

	// // eq2
	// s1 - s2 = diff
	// s1 = diff + s2

	// diff + s2 + s2 = totalSum
	// diff + 2s2 = totalSum
	// 2s2 = totalSum - diff
	// s2 = (totalSum - diff) / 2

	s2 := (totalSum - diff) / 2

	tabulation := make([][]int, n)

	for i := range tabulation {
		tabulation[i] = make([]int, s2+1)
	}

	if arr[0] == 0 {
		tabulation[0][0] = 2
	} else {
		tabulation[0][0] = 1
	}

	if arr[0] != 0 && arr[0] <= s2 {
		tabulation[0][arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := range s2 + 1 {
			pick := 0

			if arr[i] <= j {
				pick = tabulation[i-1][j-arr[i]]
			}

			notPick := tabulation[i-1][j]

			tabulation[i][j] = pick + notPick
		}
	}

	return tabulation[n-1][s2]
}

func countPartitionsOptimized(n int, diff int, arr []int) int {
	totalSum := 0

	for i := range arr {
		totalSum += arr[i]
	}

	if (totalSum-diff) < 0 || (totalSum-diff)%2 == 1 {
		return 0
	}

	// // eq1
	// s1 + s2 = totalSum

	// // eq2
	// s1 - s2 = diff
	// s1 = diff + s2

	// diff + s2 + s2 = totalSum
	// diff + 2s2 = totalSum
	// 2s2 = totalSum - diff
	// s2 = (totalSum - diff) / 2

	s2 := (totalSum - diff) / 2

	previous := make([]int, s2+1)

	if arr[0] == 0 {
		previous[0] = 2
	} else {
		previous[0] = 1
	}

	if arr[0] != 0 && arr[0] <= s2 {
		previous[arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		current := make([]int, s2+1)

		for j := range s2 + 1 {
			pick := 0

			if arr[i] <= j {
				pick = previous[j-arr[i]]
			}

			notPick := previous[j]

			current[j] = pick + notPick
		}

		previous = append([]int{}, current...)
	}

	return previous[s2]
}

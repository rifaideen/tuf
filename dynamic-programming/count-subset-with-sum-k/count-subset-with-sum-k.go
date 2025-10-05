package main

func perfectSum(arr []int, K int) int {
	n := len(arr)
	memoized := make([][]int, n)

	for i := range memoized {
		memoized[i] = make([]int, K+1)

		for j := range memoized[i] {
			memoized[i][j] = -1
		}
	}

	return f(n-1, K, arr, &memoized)
}

func f(index, target int, nums []int, memoized *[][]int) int {
	if target == 0 {
		return 1
	}

	if index == 0 {
		if nums[0] == target {
			return 1
		}

		return 0
	}

	if (*memoized)[index][target] != -1 {
		return (*memoized)[index][target]
	}

	take := 0

	if nums[index] <= target {
		take = f(index-1, target-nums[index], nums, memoized)
	}

	notTake := f(index-1, target, nums, memoized)

	(*memoized)[index][target] = take + notTake

	return (*memoized)[index][target]
}

func perfectSumT(arr []int, K int) int {
	n := len(arr)
	tabulation := make([][]int, n)

	for i := range tabulation {
		tabulation[i] = make([]int, K+1)

		tabulation[i][0] = 1
	}

	if arr[0] <= K {
		tabulation[0][arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		for sum := 1; sum <= K; sum++ {
			pick := 0

			if arr[i] <= sum {
				pick = tabulation[i-1][sum-arr[i]]
			}

			notPick := tabulation[i-1][sum]

			tabulation[i][sum] = pick + notPick
		}
	}

	return tabulation[n-1][K]
}

func perfectSumOptimized(arr []int, K int) int {
	n := len(arr)
	previous := make([]int, K+1)

	previous[0] = 1

	if arr[0] <= K {
		previous[arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		current := make([]int, K+1)
		current[0] = 1

		for sum := 1; sum <= K; sum++ {
			pick := 0

			if arr[i] <= sum {
				pick = previous[sum-arr[i]]
			}

			notPick := previous[sum]

			current[sum] = pick + notPick
		}

		previous = append([]int{}, current...)
	}

	return previous[K]
}

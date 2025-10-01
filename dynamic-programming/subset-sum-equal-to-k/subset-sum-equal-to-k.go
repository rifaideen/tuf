package main

func isSubsetSum(arr []int, target int) bool {
	n := len(arr)

	memoization := make([][]int, n)

	for i := range memoization {
		memoization[i] = make([]int, target+1)

		for j := range memoization[i] {
			memoization[i][j] = -1
		}
	}

	return f(n-1, target, arr, &memoization)
}

func f(index, target int, arr []int, memoization *[][]int) bool {
	if target == 0 {
		return true
	}

	if index == 0 {
		return arr[0] == target
	}

	if (*memoization)[index][target] != -1 {
		return (*memoization)[index][target] == 1
	}

	take := false

	if arr[index] <= target {
		take = f(index-1, target-arr[index], arr, memoization)
	}

	notTake := f(index-1, target, arr, memoization)

	if take || notTake {
		(*memoization)[index][target] = 1
	} else {
		(*memoization)[index][target] = 0
	}

	return (*memoization)[index][target] == 1
}

func isSubsetSumT(arr []int, target int) bool {
	n := len(arr)

	// DP table: tabulation[i][j] = true if we can achieve sum 'j'
	// using only the first (i+1) elements (from index 0 to i)
	//
	// Why this structure?
	// - We build solutions bottom-up: small sets → full set
	// - Each row adds one more element, expanding what sums we can form
	// - Final answer must consider ALL elements → last row (n-1)
	tabulation := make([][]bool, n)

	for i := range tabulation {
		tabulation[i] = make([]bool, target+1)

		// Base case: sum = 0 is always achievable by picking *no elements*
		// This holds for every prefix of the array
		tabulation[i][0] = true
	}

	k := target // Just to clarify usage; k is our target sum

	// Handle first element separately:
	// With only arr[0], we can only form sum = 0 (already set) and sum = arr[0]
	if arr[0] <= k {
		tabulation[0][arr[0]] = true
	}

	// Fill the table row by row: adding one number at a time
	for i := 1; i < n; i++ {
		for currSum := 1; currSum <= k; currSum++ {
			pick := false

			// If current number fits into currSum, try including it
			// Then check if the *remaining sum* (currSum - arr[i]) was achievable
			// without this element (i.e., in previous row)
			if arr[i] <= currSum {
				pick = tabulation[i-1][currSum-arr[i]]
			}

			// Not picking arr[i]: then currSum must have been achievable
			// using only the first i elements (same sum, previous row)
			notPick := tabulation[i-1][currSum]

			// Current state is valid if *either* choice works
			tabulation[i][currSum] = pick || notPick
		}
	}

	// Why return tabulation[n-1][k]?
	//
	// - Row n-1 means we've considered *all* elements in the array.
	//   Earlier rows only use subsets of the input — incomplete information.
	//
	// - Column k corresponds to the exact target sum we care about.
	//
	// - The process builds up possible sums incrementally. Only at the last row
	//   do we know whether the *entire* set can produce the target.
	//
	// Thus, tabulation[n-1][k] represents:
	// "Can we form sum = target using any subset of the full array?"
	//
	// That’s precisely the definition of the subset sum problem.
	return tabulation[n-1][k]
}

func isSubsetSumOptimized(arr []int, target int) bool {
	n := len(arr)

	// We only need two rows at a time:
	// - 'previous': results from processing first i elements
	// - 'current': building result for first i+1 elements
	//
	// Why? Because in the DP recurrence:
	//   dp[i][sum] = dp[i-1][sum] OR dp[i-1][sum - arr[i]]
	// It only depends on the previous row, NOT earlier ones.
	// So we can discard all prior rows → space optimization!
	previous := make([]bool, target+1)
	current := make([]bool, target+1)

	// Base case: sum 0 is always achievable (by picking no elements)
	previous[0] = true
	current[0] = true // maintain consistency

	// Initialize base case for first element
	if arr[0] <= target {
		previous[arr[0]] = true // with only arr[0], we can form sum = arr[0]
	}

	// Process each element one by one
	for i := 1; i < n; i++ {
		for sum := 1; sum <= target; sum++ {
			pick := false

			// If current number fits into 'sum', try including it
			// Then check if the remaining amount (sum - arr[i]) was achievable
			// using the previous state (i.e., without current element)
			if arr[i] <= sum {
				pick = previous[sum-arr[i]]
			}

			// Not picking current element:
			// same sum must have been possible without it
			notPick := previous[sum]

			// Current state: either choice works
			current[sum] = pick || notPick
		}

		// After finishing row i, 'current' becomes the new 'previous'
		// for the next iteration. We don't need older data anymore.
		//
		// This is the key to space optimization:
		// Instead of keeping all n rows, we keep just the last one.
		previous = append([]bool{}, current...)
	}

	return previous[target]
}

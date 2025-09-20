package main

// // using memoization
// func rob(nums []int) int {
// 	n := len(nums)
// 	memoized := make([]int, n)

// 	// Initialize memo table with -1 to indicate uncomputed states
// 	for i := range memoized {
// 		memoized[i] = -1
// 	}

// 	// Start from the last house and compute maximum robbery amount
// 	return f(n-1, nums, &memoized)
// }

// // f returns max money that can be robbed from houses [0..i]
// func f(i int, nums []int, memoized *[]int) int {
// 	if i == 0 {
// 		return nums[0] // Only one house available
// 	}

// 	if i < 0 {
// 		return 0 // No houses to rob
// 	}

// 	if (*memoized)[i] != -1 {
// 		return (*memoized)[i] // Return cached result if already computed
// 	}

// 	// Option 1: Rob current house + max from two houses before
// 	pick := nums[i] + f(i-2, nums, memoized)
// 	// Option 2: Skip current house, take max from previous house
// 	nonPick := f(i-1, nums, memoized)

// 	// Store the best choice for future reference
// 	(*memoized)[i] = max(pick, nonPick)

// 	return (*memoized)[i]
// }

// // using tabulation
// func rob(nums []int) int {
// 	n := len(nums)
// 	tabulation := make([]int, n)

// 	if n == 1 {
// 		return nums[0]
// 	}

// 	if n == 2 {
// 		return max(nums[0], nums[1])
// 	}

// 	// Tabulation array to store the maximum money
// 	// that can be robbed from houses [0..i] for each i
// 	for i := range tabulation {
// 		tabulation[i] = -1
// 	}

// 	// Fill the table iteratively using dynamic programming
// 	return f(n, nums, tabulation)
// }

// // f computes the maximum amount of money that can be robbed
// // by building the solution from left to right (bottom-up).
// // At each house i:
// //   - If we rob house i, we add its value to the result from house i-2
// //   - If we skip house i, we take the result from house i-1
// //
// // We choose the maximum of these two options and store it.
// func f(n int, nums []int, tabulation []int) int {
// 	tabulation[0] = nums[0] // Max profit from first house is its own value

// 	for i := 1; i < n; i++ {
// 		// Rob current house: take its value + max profit up to house i-2
// 		pick := nums[i] + 0
// 		if i-2 >= 0 {
// 			pick = nums[i] + tabulation[i-2]
// 		}

// 		// Skip current house: carry forward max profit up to house i-1
// 		nonPick := tabulation[i-1]

// 		// Store the best choice at house i
// 		tabulation[i] = max(pick, nonPick)
// 	}

// 	return tabulation[n-1] // Max profit up to the last house
// }

// using space optimization
func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	// Space-optimized iterative solution using two variables instead of a full DP array
	// We only need the results from the last two houses at any step
	return f(n, nums)
}

// f computes the maximum amount of money that can be robbed
// by building the solution from left to right (bottom-up).
// At each house i:
//   - If we rob house i, we add its value to the result from house i-2
//   - If we skip house i, we take the result from house i-1
//
// We choose the maximum of these two options.
// Instead of storing all states, we track only:
//   - i1: max profit up to the previous house (i-1)
//   - i2: max profit up to the house before that (i-2)
//
// This reduces space complexity from O(n) to O(1).
func f(n int, nums []int) int {
	i1 := nums[0] // max profit up to house 0
	i2 := 0       // max profit up to house -1 (assumed 0)

	for i := 1; i < n; i++ {
		pick := nums[i] // rob current house
		if i > 1 {
			pick += i2 // add profit from house i-2 if exists
		}

		nonPick := i1 // skip current house → profit = best up to i-1

		current := max(pick, nonPick) // best choice at house i

		// update for next iteration: shift window forward
		i2 = i1      // i2 becomes the result from one step back
		i1 = current // i1 becomes current result
	}

	return i1 // final answer: max profit up to last house
}

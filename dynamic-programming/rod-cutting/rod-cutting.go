package main

// RodCutting solves the rod cutting problem using recursion.
// Returns the maximum profit obtainable by cutting a rod of length n,
// given the price for each possible piece length (1 to n).
func RodCutting(price []int, n int) int {
	return f(n-1, n, price) // Start recursion from last index and target = full rod length
}

// f(index, target, price): Recursive helper function
// - index: current index in price array (represents rod piece length = index+1)
// - target: remaining rod length to cut
// - We try two choices: pick a piece of length (index+1), or skip it
func f(index, target int, price []int) int {
	// Base case: if we're at the first piece (length 1), we can take 'target' pieces of it
	if index == 0 {
		return price[0] * target // Only option: use length-1 pieces
	}

	pick := 0
	rodLength := index + 1 // The actual length of the rod piece represented by this index

	// If the current rod piece fits into the remaining length, we can choose to pick it
	if rodLength <= target {
		// Take one piece of current length -> add its price and recurse on reduced target
		// Note: we stay at same index because we can reuse the same length (unbounded choice)
		pick = price[index] + f(index, target-rodLength, price)
	}

	// Option to not take the current piece -> move to next smaller piece
	notPick := f(index-1, target, price)

	// Return maximum profit from either choice
	return max(pick, notPick)
}

// RodCuttingMemoized: Top-down DP with memoization to avoid recomputation
func RodCuttingMemoized(price []int, n int) int {
	memoized := make([][]int, n)

	// Initialize memo table: memo[i][t] stores max profit for rod of length t using pieces up to index i
	for i := range memoized {
		memoized[i] = make([]int, n+1)
		for j := range memoized[i] {
			memoized[i][j] = -1 // -1 indicates not computed yet
		}
	}

	return fMemoized(n-1, n, price, &memoized)
}

// fMemoized: Memoized version of recursive function
func fMemoized(index, target int, price []int, memoized *[][]int) int {
	// Base case: only length-1 pieces available
	if index == 0 {
		return price[0] * target
	}

	// If already computed, return stored result
	if (*memoized)[index][target] != -1 {
		return (*memoized)[index][target]
	}

	pick := 0
	rodLength := index + 1

	// If current piece fits, consider picking it (can reuse same length)
	if rodLength <= target {
		pick = price[index] + fMemoized(index, target-rodLength, price, memoized)
	}

	// Option to skip current piece
	notPick := fMemoized(index-1, target, price, memoized)

	// Store and return best result
	(*memoized)[index][target] = max(pick, notPick)
	return (*memoized)[index][target]
}

// RodCuttingTabulation: Bottom-up DP (tabulation)
// Builds solution iteratively without recursion
func RodCuttingTabulation(price []int, n int) int {
	// dp[i][t] = max profit using pieces up to length (i+1) for a rod of length t
	tabulation := make([][]int, n)
	for i := range tabulation {
		tabulation[i] = make([]int, n+1) // n+1 lengths: 0 to n
	}

	// Base case: first row — only pieces of length 1 are allowed
	for target := 0; target <= n; target++ {
		tabulation[0][target] = price[0] * target // Use 'target' number of length-1 pieces
	}

	// Fill the table row by row
	for i := 1; i < n; i++ {
		for target := 0; target <= n; target++ {
			pick := 0
			rodLength := i + 1 // Current piece length under consideration

			// If current piece fits, we can pick it
			if rodLength <= target {
				// Add price of one piece of this length and use precomputed result
				// Note: we stay in same row because multiple uses allowed
				pick = price[i] + tabulation[i][target-rodLength]
			}

			// Skip current piece: best result using only previous pieces
			notPick := tabulation[i-1][target]

			// Max of picking or skipping
			tabulation[i][target] = max(pick, notPick)
		}
	}

	// Answer: max profit using all piece types (0 to n-1) for rod of length n
	return tabulation[n-1][n]
}

// RodCuttingOptimized: Space-optimized version
// Instead of storing entire table, keep only previous row
func RodCuttingOptimized(price []int, n int) int {
	// previous[t] = max profit for rod length t using pieces considered so far
	previous := make([]int, n+1)

	// Initialize for first piece (length 1)
	for target := 0; target <= n; target++ {
		previous[target] = price[0] * target
	}

	// Process each piece length from 2 to n
	for i := 1; i < n; i++ {
		current := make([]int, n+1) // current state for including up to piece length (i+1)

		for target := 0; target <= n; target++ {
			pick := 0
			rodLength := i + 1

			// If we can use this piece, do so and add its value
			if rodLength <= target {
				// Reuse same row (current) since unlimited cuts allowed
				pick = price[i] + current[target-rodLength]
			}

			// Not picking: rely on results from previous set of pieces
			notPick := previous[target]

			current[target] = max(pick, notPick)
		}

		// Move to next iteration: current becomes previous
		previous = append([]int{}, current...) // Deep copy
	}

	// Final answer after considering all pieces
	return previous[n]
}

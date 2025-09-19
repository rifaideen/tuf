package main

import "math"

// // frogJump: Pure recursive solution (no memoization).
// // Computes min energy to reach last stone from stone 0.
// // Only for understanding — too slow for large inputs.
// func frogJump(heights []int, k int) int {
//     return f(len(heights)-1, k, heights) // start from last stone
// }

// // f: returns min energy to reach stone i from stone 0.
// // Base case: i == 0 → cost 0.
// // Otherwise: try all k possible jumps back, pick the cheapest path.
// // Warning: recalculates same states repeatedly → exponential time.
// func f(i, k int, heights []int) int {
//     if i == 0 {
//         return 0 // already at start
//     }

//     minEnergy := math.MaxInt

//     // try jumping from up to k steps back
//     for j := 1; j <= k; j++ {
//         if i-j >= 0 {
//             // cost = cost to reach (i-j) + jump cost to i
//             energy := f(i-j, k, heights) + int(math.Abs(float64(heights[i] - heights[i-j])))
//             minEnergy = min(minEnergy, energy)
//         }
//     }

//     return minEnergy // return best option
// }

// frogJump: Top-down recursive DP with memoization.
// Starts from the last stone and computes minimum energy to reach it from stone 0.
// func frogJump(heights []int, k int) int {
//     // memoized[i] stores min energy to reach stone i from stone 0
//     // initialize all values as -1 (uncomputed)
//     memoized := make([]int, len(heights))
//     for i := range memoized {
//         memoized[i] = -1
//     }

//     // start recursion from the last stone (index = len-1)
//     return f(len(heights)-1, k, heights, &memoized)
// }

// // f: returns min energy to reach stone i from stone 0
// // Base case: stone 0 → cost 0
// // Memo check: if already computed, return cached result
// // Otherwise: try jumping from up to k steps back, pick the cheapest path
// func f(i, k int, heights []int, memoized *[]int) int {
//     if i == 0 {  // base case: at starting stone
//         return 0
//     }

//     if (*memoized)[i] != -1 {  // already computed
//         return (*memoized)[i]
//     }

//     minEnergy := math.MaxInt  // track minimum energy path

//     // try all possible jumps: 1 to k steps back
//     for j := 1; j <= k; j++ {
//         if i-j >= 0 {  // valid previous stone
//             // total energy = energy to reach (i-j) + jump cost to i
//             energy := f(i-j, k, heights, memoized) + int(math.Abs(float64(heights[i] - heights[i-j])))
//             minEnergy = min(minEnergy, energy)
//         }
//     }

//     (*memoized)[i] = minEnergy  // cache result before returning
//     return minEnergy
// }

// frogJump: Top-down DP with memoization.
// Starts from last stone and recurses down to 0.
// func frogJump(heights []int, k int) int {
// 	// memoized[i] = min energy to reach last stone from stone i
// 	// initialized to -1 (uncomputed)
// 	memoized := make([]int, len(heights))
// 	for i := range memoized {
// 		memoized[i] = -1
// 	}

// 	// Start recursion from the last stone (index len-1)
// 	return f(len(heights)-1, k, heights, &memoized)
// }

// // f: Returns min energy to reach stone i from stone 0 (via reverse recursion).
// // Uses memoization to avoid recomputing states.
// func f(i, k int, heights []int, memoized *[]int) int {
// 	if i == 0 {
// 		return 0 // base case: already at start
// 	}

// 	if (*memoized)[i] != -1 {
// 		return (*memoized)[i] // return cached result
// 	}

// 	minEnergy := math.MaxInt

// 	// try all jumps of 1 to k steps back
// 	for j := 1; j <= k; j++ {
// 		if i-j >= 0 {
// 			// cost = cost to reach i-j + jump cost from i-j to i
// 			energy := f(i-j, k, heights, memoized) + int(math.Abs(float64(heights[i]-heights[i-j])))
// 			minEnergy = min(minEnergy, energy)
// 		}
// 	}

// 	(*memoized)[i] = minEnergy // cache result
// 	return minEnergy
// }

// frogJump: Returns min cost for frog to reach last stone.
// Mimics TUF's Java version: uses bottom-up DP (tabulation).
func frogJump(heights []int, k int) int {
	// dp array: tabulation[i] = min cost to reach stone i
	tabulation := make([]int, len(heights))

	// Pass total count (n = len) to f, not last index
	return f(len(heights), k, heights, &tabulation)
}

// f: Tabulation function.
// Base case: stone 0 has cost 0.
// For each stone i from 1 to n-1:
//
//	try jumping from up to k steps back
//	dp[i] = min(dp[i-j] + |height[i]-height[i-j]|)
//
// Return dp[n-1] → cost to reach last stone.
func f(n, k int, heights []int, tabulation *[]int) int {
	(*tabulation)[0] = 0 // start at first stone, zero cost

	// fill dp table for stones 1 to n-1 (i < n)
	for i := 1; i < n; i++ {
		minSteps := math.MaxInt

		// check all possible previous stones within k range
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				// cost = prev cost + jump energy
				jumps := (*tabulation)[i-j] + int(math.Abs(float64(heights[i]-heights[i-j])))
				minSteps = min(minSteps, jumps)
			}
		}

		(*tabulation)[i] = minSteps // store best cost for stone i
	}

	return (*tabulation)[n-1] // final answer: cost to reach last stone
}

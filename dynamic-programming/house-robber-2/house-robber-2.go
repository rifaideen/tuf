package main

func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	// Special handling for circular arrangement: house[0] and house[n-1] are adjacent.
	// So we can't rob both. We break the circle by considering two cases:
	// Case 1: Rob houses from index 0 to n-2 (exclude last house)
	// Case 2: Rob houses from index 1 to n-1 (exclude first house)
	// Answer is the maximum of these two cases.
	return max(f(n-1, nums[0:n-1]), f(n-1, nums[1:]))
}

// f computes the maximum amount that can be robbed in a linear row of houses (non-circular).
// Uses space-optimized DP where:
// i1 = max profit up to previous house (i-1)
// i2 = max profit up to house before that (i-2)
// At each house, choose between:
// - Rob it: current value + i2 (best from two houses back)
// - Skip it: take i1 (best from previous house)
// Then update i2 and i1 to move the window forward.
func f(n int, nums []int) int {
	i1 := nums[0] // max profit up to the first house
	i2 := 0       // max profit before the first house (base case)

	for i := 1; i < n; i++ {
		pick := nums[i] // rob current house
		if i > 1 {
			pick += i2 // add profit from house i-2 if exists
		}

		nonPick := i1 // skip current house → profit = best up to i-1

		current := max(pick, nonPick) // best choice at house i

		// slide window forward
		i2 = i1      // previous result becomes "two steps back"
		i1 = current // current result becomes "one step back"
	}

	return i1 // maximum profit for this linear segment
}

package main

func canPartition(nums []int) bool {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	// If the total sum is odd, we cannot split into two equal subsets
	if sum%2 != 0 {
		return false
	}

	// We now need to find a subset that sums up to half of total sum
	return subsetSumTarget(nums, sum/2)
}

func subsetSumTarget(nums []int, target int) bool {
	n := len(nums)

	// Space-optimized DP: only keep track of previous row and current row
	previous := make([]bool, target+1) // dp values from the previous row
	current := make([]bool, target+1)  // dp values for the current row

	// Base case: sum of 0 is always possible (by selecting no elements)
	previous[0] = true
	current[0] = true

	// Initialize the first row: with only the first element, we can only form sum = nums[0]
	if nums[0] <= target {
		previous[nums[0]] = true
	}

	// Fill the table row by row (i.e., element by element)
	for i := 1; i < n; i++ {
		for sum := 1; sum <= target; sum++ {
			pick := false

			// We can only pick the current number if it doesn't exceed the current sum
			if nums[i] <= sum {
				// Check if the remaining sum (sum - nums[i]) was achievable using elements before current
				pick = previous[sum-nums[i]]
			}

			// If we don't pick the current number, then the sum must have been achievable using previous elements
			notPick := previous[sum]

			// Current state is true if either picking or not picking leads to a valid solution
			current[sum] = pick || notPick
		}

		// After processing current row, copy its values to 'previous' for next iteration
		// This avoids referencing issues and simulates moving to the next row
		previous = append([]bool{}, current...)
	}

	// The final answer is whether we can achieve the target sum using all elements
	return previous[target]
}

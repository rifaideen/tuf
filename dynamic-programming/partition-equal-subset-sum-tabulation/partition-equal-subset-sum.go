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

	// Create a 2D DP table: n x (target + 1), all initialized to false
	tabulation := make([][]bool, n)

	for i := range tabulation {
		tabulation[i] = make([]bool, target+1)

		// Base case: a sum of 0 is always possible (by picking no elements)
		tabulation[i][0] = true
	}

	// For the first element, we can only achieve the sum equal to that element
	if nums[0] <= target {
		tabulation[0][nums[0]] = true
	}

	// Fill the DP table bottom-up
	for i := 1; i < n; i++ {
		for sum := 1; sum <= target; sum++ {
			pick := false

			// We can only pick current number if it doesn't exceed the current sum
			if nums[i] <= sum {
				// If we pick nums[i], check if the remaining sum (sum - nums[i]) was achievable without it
				pick = tabulation[i-1][sum-nums[i]]
			}

			// If we don't pick nums[i], then the sum must have been achievable using previous elements
			notPick := tabulation[i-1][sum]

			// Current cell is true if either picking or not picking leads to a valid solution
			tabulation[i][sum] = pick || notPick
		}
	}

	// The final answer is whether we can achieve 'target' using all elements
	return tabulation[n-1][target]
}

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

	// Create a 2D memoization table: n x (target + 1), initialized to -1
	memoized := make([][]int, n)

	for i := range memoized {
		memoized[i] = make([]int, target+1)

		for j := range memoized[i] {
			// -1 indicates this state hasn't been computed yet
			memoized[i][j] = -1
		}
	}

	// Start recursion from the last index with desired target
	return f(n-1, target, nums, &memoized)
}

func f(index, target int, nums []int, memoized *[][]int) bool {
	// Base case: if target is achieved, we found a valid subset
	if target == 0 {
		return true
	}

	// Base case: if we're at the first element, check if it matches the remaining target
	if index == 0 {
		return nums[0] == target
	}

	// Check if result for current index and target is already computed
	if (*memoized)[index][target] != -1 {
		// Return true if previously computed result was 1 (true), else false
		return (*memoized)[index][target] == 1
	}

	// Try picking the current element, but only if it doesn't exceed the target
	pick := false
	if nums[index] <= target {
		// Reduce target by nums[index] and move to previous index
		pick = f(index-1, target-nums[index], nums, memoized)
	}

	// Try not picking the current element, carry same target to previous index
	notPick := f(index-1, target, nums, memoized)

	// If either picking or not picking leads to a solution, mark as success (1)
	if pick || notPick {
		(*memoized)[index][target] = 1
	} else {
		// Otherwise, mark as failure (0)
		(*memoized)[index][target] = 0
	}

	// Return whether a valid subset was found for this state
	return (*memoized)[index][target] == 1
}

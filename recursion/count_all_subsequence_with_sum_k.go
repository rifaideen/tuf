package main

func countSubsequenceWithTargetSum(nums []int, k int) int {
	n := len(nums)

	return countSubsequenceWithTargetRecursive(0, 0, nums, k, n)
}

func countSubsequenceWithTargetRecursive(index, sum int, nums []int, k, n int) int {
	// base case satified
	if index == n {
		// target sum satisfied
		if sum == k {
			return 1
		}

		// target sum not satisfied
		return 0
	}

	// pick the item at current index
	left := countSubsequenceWithTargetRecursive(index+1, sum+nums[index], nums, k, n)

	// don't pick the item at current index
	right := countSubsequenceWithTargetRecursive(index+1, sum, nums, k, n)

	return left + right
}

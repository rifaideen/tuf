package main

func CheckSubsequenceSum(nums []int, k int) bool {
	n := len(nums)

	return checkSubsequenceSumRecursively(0, 0, nums, k, n)
}

func checkSubsequenceSumRecursively(index int, sum int, nums []int, k, n int) bool {
	// base condition satisfied
	if index == n {
		// check if constraint satisfied
		return sum == k
	}

	// pick
	if checkSubsequenceSumRecursively(index+1, sum+nums[index], nums, k, n) {
		return true
	}

	// don't pick
	if checkSubsequenceSumRecursively(index+1, sum, nums, k, n) {
		return true
	}

	return false
}

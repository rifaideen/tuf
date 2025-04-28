package main

import "slices"

func subsetSum(nums []int) []int {
	ans := []int{}

	subsetSumRecursive(0, len(nums), 0, &ans, nums)

	// since the problem wants the answer in the sorted order, we sort it.
	slices.Sort(ans)

	return ans
}

func subsetSumRecursive(index, n, sum int, ans *[]int, nums []int) {
	if index == n {
		*ans = append(*ans, sum)
		return
	}

	// pick the current element and move next
	subsetSumRecursive(index+1, n, sum+nums[index], ans, nums)

	// don't pick the current element and move next
	subsetSumRecursive(index+1, n, sum, ans, nums)
}

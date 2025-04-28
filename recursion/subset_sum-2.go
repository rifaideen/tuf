package main

import "slices"

func subsetSumTwo(nums []int) [][]int {
	ans := [][]int{}
	slices.Sort(nums)

	subsetSumTwoRecursive(0, len(nums), &ans, &[]int{}, nums)

	return ans
}

func subsetSumTwoRecursive(index, n int, ans *[][]int, sequence *[]int, nums []int) {
	tmp := make([]int, len(*sequence))
	copy(tmp, *sequence)
	*ans = append(*ans, tmp)

	for i := index; i < n; i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		// pick it
		*sequence = append(*sequence, nums[i])
		subsetSumTwoRecursive(i+1, n, ans, sequence, nums)
		*sequence = (*sequence)[:len(*sequence)-1]
	}
}

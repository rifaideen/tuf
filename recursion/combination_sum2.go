package main

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	// we need the answer to be in the sorted manner, so sorting the candiates will help us create the
	// sorted answer
	slices.Sort(candidates)
	ans := [][]int{}

	combinationSumRecursive(0, &ans, &[]int{}, candidates, target)

	return ans
}

func combinationSumRecursive(index int, ans *[][]int, sequence *[]int, candidates []int, target int) {
	// at any point of time, we reach the target 0. then add them to the answers array
	if target == 0 {
		tmp := make([]int, len(*sequence))
		copy(tmp, *sequence)
		*ans = append(*ans, tmp)

		return
	}

	for i := index; i < len(candidates); i++ {
		// to avoid duplicates and still allow us the process the first index for the first time
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		// we can't continue when the cadidate value in greater than target, hence break it
		if candidates[i] > target {
			break
		}

		// pick up the current candidate
		*sequence = append(*sequence, candidates[i])
		combinationSumRecursive(i+1, ans, sequence, candidates, target-candidates[i])

		// take it back after the iteration, next iteration will be adding it
		*sequence = (*sequence)[:len(*sequence)-1]
	}
}

package main

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := [][]int{}

	// we start with index = 0 and empty sequence and pass the sequence and answer by reference
	combinationSumRecusrive(0, target, n, &ans, &[]int{}, candidates)

	return ans
}

func combinationSumRecusrive(index, target, n int, ans *[][]int, sequence *[]int, candidates []int) {
	if index == n {
		if target == 0 {
			// create the copy of sequence before appending to it
			// since the sequence is pointer, appending them would cause side-effects when the sequence got updated in other places.
			tmp := make([]int, len(*sequence))
			copy(tmp, *sequence)

			*ans = append(*ans, tmp)
		}

		return
	}

	// Since we are allowed to pickup the same element multiple times, we slightly modify the approach
	//
	// try to pickup the same element as long as the current value is <= taget value and decrease the target with current value
	//
	// pickup until the current value is <= target value
	if candidates[index] <= target {
		// append the current value and pass
		*sequence = append(*sequence, candidates[index])

		// recursively check how many times we can pickup
		combinationSumRecusrive(index, target-candidates[index], n, ans, sequence, candidates)

		// pullback the last element added
		*sequence = (*sequence)[:len(*sequence)-1]
	}

	// don't pick and move on
	combinationSumRecusrive(index+1, target, n, ans, sequence, candidates)
}

package main

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}

	// to find all the possible subsets, it will take upto (2^n)-1 combinations
	//
	// shifing 1 to the left n times gives 2^n
	for i := 0; i < (1 << n); i++ {
		// create sub-set to store the combination
		subset := []int{}

		// loop j till n times and check if j-th bit is set in i
		for j := 0; j < n; j++ {
			// using bitwise and operation to see
			if (i & (1 << j)) != 0 {
				subset = append(subset, nums[j])
			}
		}

		ans = append(ans, subset)
	}

	return ans
}

package main

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}

	for i := 0; i < (1 << n); i++ {
		subset := []int{}

		for j := 0; j < n; j++ {
			if (i & (1 << i)) != 0 {
				subset = append(subset, nums[j])
			}
		}

		ans = append(ans, subset)
	}

	return ans
}

package main

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	seq := []int{}

	combinationSum3Recursive(0, &seq, k, n, &ans)

	return ans
}

func combinationSum3Recursive(index int, sequence *[]int, k, n int, ans *[][]int) {
	// base case: target becomes 0 and sequence length = k
	if n == 0 && len(*sequence) == k {
		tmp := make([]int, len(*sequence))
		copy(tmp, *sequence)
		*ans = append(*ans, tmp)

		return
	} else if n < 0 || len(*sequence) == k {
		// target becomes negative or sequence length = k
		return
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := index; i < 9; i++ {
		// we don't need to process further when nums[i] is greater than target
		if nums[i] > n {
			break
		}

		// pick it
		*sequence = append(*sequence, nums[i])
		combinationSum3Recursive(i+1, sequence, k, n-nums[i], ans)

		// pull it back, so in the next iteration it will act don't pick
		*sequence = (*sequence)[:len(*sequence)-1]
	}
}

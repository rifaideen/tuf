package main

func isSubsetSum(arr []int, target int) bool {
	n := len(arr)

	memoization := make([][]int, n)

	for i := range memoization {
		memoization[i] = make([]int, target+1)

		for j := range memoization[i] {
			memoization[i][j] = -1
		}
	}

	return f(n-1, target, arr, &memoization)
}

func f(index, target int, arr []int, memoization *[][]int) bool {
	if target == 0 {
		return true
	}

	if index == 0 {
		return arr[0] == target
	}

	if (*memoization)[index][target] != -1 {
		return (*memoization)[index][target] == 1
	}

	take := false

	if arr[index] <= target {
		take = f(index-1, target-arr[index], arr, memoization)
	}

	notTake := f(index-1, target, arr, memoization)

	if take || notTake {
		(*memoization)[index][target] = 1
	} else {
		(*memoization)[index][target] = 0
	}

	return (*memoization)[index][target] == 1
}

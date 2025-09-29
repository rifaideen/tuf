package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memoization := make([][]int, n)

	for i := range memoization {
		memoization[i] = make([]int, n)

		for j := range memoization[i] {
			memoization[i][j] = -1e9
		}
	}

	return f(0, 0, n, triangle, &memoization)
}

func f(i, j, n int, triangle [][]int, memoization *[][]int) int {
	if i == (n - 1) {
		return triangle[i][j]
	}

	if (*memoization)[i][j] != -1e9 {
		return (*memoization)[i][j]
	}

	down := f(i+1, j, n, triangle, memoization)
	downRight := f(i+1, j+1, n, triangle, memoization)

	(*memoization)[i][j] = triangle[i][j] + min(down, downRight)

	return (*memoization)[i][j]
}

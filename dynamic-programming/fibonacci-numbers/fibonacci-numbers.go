package main

// fibonacci numbers using memoization
// T.C: O(n)
// S.C: O(n) + O(n) // one for recursive stack and another for memoization array
func fibonacci(n int, memoized *[]int) int {
	if n <= 1 {
		return n
	}

	if (*memoized)[n] != -1 {
		return (*memoized)[n]
	}

	(*memoized)[n] = fibonacci(n-1, memoized) + fibonacci(n-2, memoized)

	return (*memoized)[n]
}

// fibonacci numbers using tabulation
// T.C: O(n)
// S.C: O(n) // for tabulation array and we have successfully get rid of recursive stack space
// by using tabulization
func fibonacci2(n int) int {
	tabular := make([]int, n+1)
	tabular[0] = 0
	tabular[1] = 1

	for i := 2; i <= n; i++ {
		tabular[i] = tabular[i-1] + tabular[i-2]
	}

	return tabular[n]
}

// fibonacci numbers using couple variables in a pattern
// T.C: O(n)
// S.C: O(1) // we have entirely optimized the space we used by having the previous, second previous and current variables
// by using tabulization
func fibonacci3(n int) int {
	previous := 1
	secondPrevious := 0

	for i := 2; i <= n; i++ {
		current := previous + secondPrevious
		secondPrevious = previous
		previous = current
	}

	return previous
}

package main

import (
	"math"
)

// naive recursive solution
func frogJump(heights []int) int {
	return fibonacci(len(heights)-1, heights)
}

func fibonacci(n int, heights []int) int {
	if n == 0 {
		return 0
	}

	left := fibonacci(n-1, heights) + int(math.Abs(float64(heights[n]-heights[n-1])))
	right := math.MaxInt

	if n > 1 {
		right = fibonacci(n-2, heights) + int(math.Abs(float64(heights[n]-heights[n-2])))
	}

	return min(left, right)
}

// using memoization solution
func frogJump2(heights []int) int {
	n := len(heights)
	memoized := make([]int, n+1)

	for i := range memoized {
		memoized[i] = -1
	}

	return fibonacci2(len(heights)-1, heights, &memoized)
}

func fibonacci2(n int, heights []int, memoized *[]int) int {
	if n == 0 {
		return 0
	}

	if (*memoized)[n] != -1 {
		return (*memoized)[n]
	}

	left := fibonacci(n-1, heights) + int(math.Abs(float64(heights[n]-heights[n-1])))
	right := math.MaxInt

	if n > 1 {
		right = fibonacci(n-2, heights) + int(math.Abs(float64(heights[n]-heights[n-2])))
	}

	(*memoized)[n] = min(left, right)

	return (*memoized)[n]
}

// using tabulation solution
func frogJump3(heights []int) int {
	n := len(heights)
	memoized := make([]int, n)

	for i := range memoized {
		memoized[i] = -1
	}

	return fibonacci3(len(heights), heights)
}

func fibonacci3(n int, heights []int) int {
	tabulation := make([]int, n)
	tabulation[0] = 0

	for i := 1; i < n; i++ {
		left := tabulation[i-1] + int(math.Abs(float64(heights[i]-heights[i-1])))
		right := math.MaxInt

		if i > 1 {
			right = tabulation[i-2] + int(math.Abs(float64(heights[i]-heights[i-2])))
		}

		tabulation[i] = min(left, right)
	}

	return tabulation[n-1]
}

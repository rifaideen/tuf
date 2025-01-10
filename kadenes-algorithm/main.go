package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindMaximumSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(FindMaximumSubArray([]int{1}))
}

func FindMaximumSubArray(nums []int) int {
	maxSum := -math.MaxInt
	sum := 0

	for _, num := range nums {
		sum += num

		if sum > maxSum {
			maxSum = sum
		}

		if sum <= 0 {
			sum = 0
		}
	}

	return maxSum
}

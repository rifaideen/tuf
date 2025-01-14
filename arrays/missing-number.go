package main

func missingSortedNumber(nums []int, n int) int {
	for i := 0; i < n-1; i++ {
		// calculating next value and comparing it with next index
		next := nums[i] + 1

		if nums[i+1] != next {
			return next
		}
	}

	return -1
}

func missingNumber(nums []int) int {
	// (n * (n+1)) / 2 to get the sum of numbers
	n := len(nums)
	s1 := (n * (n + 1)) / 2
	s2 := 0

	for _, num := range nums {
		s2 += num
	}

	// subtracting the sum of n charecters with sum of the numbers supplied
	return s1 - s2
}

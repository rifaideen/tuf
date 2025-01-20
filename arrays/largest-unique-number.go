package main

// Leetcode 1133
// Given a array of integers, find the largest integer that occurs once.
// If no integer occurs once, return -1
//
// Constraints:
// 1 <= nums.length <= 2000
// 0 <= nums[i] <= 1000
//
// Intuition:
// create new array of size 1000 (derived from the constraint) filled with -1
//
// for each iteration of nums, count the number of occurence, for first time set to 1 and increment for the next time
//
// start the iteration in reverse order and return the first occurence with value 1
func LargestUniqueNumber(nums []int) int {
	occurences := make([]int, 1000)

	// prefill array with -1
	for i := range occurences {
		occurences[i] = -1
	}

	for _, num := range nums {
		if occurences[num] == -1 {
			occurences[num] = 1
		} else {
			occurences[num]++
		}
	}

	for i := len(occurences) - 1; i >= 0; i-- {
		if occurences[i] == 1 {
			return i
		}
	}

	return -1
}

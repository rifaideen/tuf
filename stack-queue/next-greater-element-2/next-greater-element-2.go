package main

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	n := len(nums)
	ans := make([]int, n)

	// start the loop from right to left, since the array is assumed to be circular, we start the array from 2*N - 1
	for i := (2 * n) - 1; i >= 0; i-- {
		// pop all smaller elements from the stack which are less than or equal to current value at nums[i%n]
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}

		// when "i" is within n, and stack not empty, push the stack top to the answer, otherwise -1
		if i < n {
			if len(stack) > 0 {
				ans[i] = stack[len(stack)-1]
			} else {
				ans[i] = -1
			}
		}

		// push the current value at nums[i%n] to the stack
		stack = append(stack, nums[i%n])
	}

	return ans
}

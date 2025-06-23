package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	mapping := map[int]int{}

	// create a monotonic stack which keeps higher values on top
	for _, num := range nums2 {
		// while stack is not empty and stack top is less than or equl to current number
		// pop the value from stack and set the popped value and the current number to the
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			// poped value is smaller, num is it's next greater element
			mapping[stack[len(stack)-1]] = num

			stack = stack[:len(stack)-1]
		}

		// push it to stack
		stack = append(stack, num)
	}

	// pop the remaining elements in stack and set it's next greater element to -1
	for len(stack) > 0 {
		mapping[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	ans := []int{}

	// read the values from the map and push it to the answer array
	for _, num := range nums1 {
		ans = append(ans, mapping[num])
	}

	return ans
}

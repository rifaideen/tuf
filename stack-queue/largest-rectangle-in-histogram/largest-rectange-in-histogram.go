package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	pse := pse(heights)
	nse := nse(heights)
	total := 0

	for i := range n {
		total = max(total, (nse[i]-pse[i]-1)*heights[i])
	}

	return total
}

func pse(heights []int) []int {
	n := len(heights)
	pse := make([]int, n)
	stack := []int{}

	for i := range n {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			stack = stack[:len(stack)-1]
		}

		pse[i] = -1

		if len(stack) > 0 {
			pse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return pse
}

func nse(heights []int) []int {
	n := len(heights)
	nse := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		nse[i] = n

		if len(stack) > 0 {
			nse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return nse
}

// Optimized solution
func LargestRectangleAreaOptimized(heights []int) int {
	n := len(heights)
	stack := []int{}
	maxValue := 0

	// Start the loop from 0 to n-1
	for key, val := range heights {
		// calculate pse
		for len(stack) > 0 && heights[stack[len(stack)-1]] > val {
			// take the stack top and pop the stack
			element := stack[len(stack)-1]
			// pop the stack
			stack = stack[:len(stack)-1]

			// set default pse to -1 when none found
			pse := -1

			// if stack has something, that would becomes pse of current element
			if len(stack) > 0 {
				pse = stack[len(stack)-1]
			}

			// set nse to outer loop's index
			nse := key

			maxValue = max(maxValue, (nse-pse-1)*heights[element])
		}

		// stack must contain the indices, not the value
		stack = append(stack, key)
	}

	// handle the case where stack has something in it
	for len(stack) > 0 {
		// top is the current element
		element := stack[len(stack)-1]
		// pop
		stack = stack[:len(stack)-1]

		// nse is n, because stack top can't have next element
		nse := n

		// set pse to -1 by default when none found previously
		pse := -1

		// if stack has something, the top is the pse
		if len(stack) > 0 {
			pse = stack[len(stack)-1]
		}

		maxValue = max(maxValue, (nse-pse-1)*heights[element])
	}

	return maxValue
}

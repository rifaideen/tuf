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

package main

func maximalRectangle(matrix [][]byte) int {
	maxValue := 0
	rows := len(matrix)
	cols := len(matrix[0])

	// create empty prefix sum matrix
	prefixSum := make([][]int, rows)

	for row := range rows {
		prefixSum[row] = make([]int, cols)
	}

	// calculate prefix sum, when current element is 0, reset the sum to 0
	//
	// we are calculating prefix sum from top to bottom
	for col := range cols {
		sum := 0

		for row := range rows {
			sum += int(matrix[row][col] - '0')

			if int(matrix[row][col]-'0') == 0 {
				sum = 0
			}

			prefixSum[row][col] = sum
		}
	}

	for i := range rows {
		maxValue = max(maxValue, largestHistogram(prefixSum[i]))
	}

	return maxValue
}

// this is same as largest-rectangle-in-histogram
func largestHistogram(heights []int) int {
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

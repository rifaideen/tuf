package main

// Approach #1: using stack
func maxDepthStack(s string) int {
	ans := 0
	stack := []int{}

	// for each iteration, we push it to stack when '(' and pop it from stack when ')'
	// ans calculate the max length of stack to be the answer
	for _, char := range s {
		if char == '(' {
			stack = append(stack, '(')
		} else if char == ')' {
			stack = stack[0 : len(stack)-1]
		}

		ans = max(ans, len(stack))
	}

	return ans
}

// Apprach #2: using counter instead of stack
func maxDepthCounter(s string) int {
	ans := 0
	count := 0

	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}

		ans = max(ans, count)
	}

	return ans
}

package main

func removeOuterParentheses(s string) string {
	// Handle empty string case
	if s == "" {
		return s
	}

	var ans []byte
	stack := []byte{}
	start := 0

	for i := 0; i < len(s); i++ {
		// Push '(' onto the stack
		if s[i] == '(' {
			stack = append(stack, s[i])
		}

		// Pop from stack when ')' is encountered
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
		}

		// When stack becomes empty, we've found a primitive string
		if len(stack) == 0 {
			// Append characters between start+1 and i, excluding outermost parentheses
			ans = append(ans, s[start+1:i]...)
			// Update start to the next primitive string
			start = i + 1
		}
	}

	return string(ans)
}

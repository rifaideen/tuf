package main

func generateParenthesis(n int) []string {
	ans := []string{}

	var helper func(left, right, n int, text string)

	helper = func(left, right, n int, text string) {
		if left > n {
			return
		}

		if (left+right) == (2*n) && left == right {
			ans = append(ans, text)

			return
		}

		helper(left+1, right, n, text+"(")

		if left > right {
			helper(left, right+1, n, text+")")
		}
	}

	helper(0, 0, n, "")

	return ans
}

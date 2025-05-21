package main

func infix2prefix(expression string) string {
	// ans := ""

	left := 0
	n := len(expression)
	right := n - 1
	reversed := []rune(expression)

	// reverse and toggle the brackets (open to close and close to open)
	for left <= right {
		reversed[left], reversed[right] = reversed[right], reversed[left]

		if reversed[left] == '(' {
			reversed[left] = ')'
		} else if reversed[left] == ')' {
			reversed[left] = '('
		}

		if reversed[right] == '(' {
			reversed[right] = ')'
		} else if reversed[right] == ')' {
			reversed[right] = '('
		}

		left++
		right--
	}

	reversed = []rune("")

	left = 0
	right = len(reversed) - 1

	for left <= right {
		reversed[left], reversed[right] = reversed[right], reversed[left]

		if reversed[left] == '(' {
			reversed[left] = ')'
		} else if reversed[left] == ')' {
			reversed[left] = '('
		}

		if reversed[right] == '(' {
			reversed[right] = ')'
		} else if reversed[right] == ')' {
			reversed[right] = '('
		}

		left++
		right--
	}

	return string(reversed)
}

// func infix2postfix(expression string) string {
// 	stack := arraystack.New()
// 	ans := ""

// 	for _, char := range len(expression) {
// 		if char >= 'A' && char <= 'Z'
// 	}

// 	return ans
// }

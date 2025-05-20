package main

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

func infix2postfix(expr string) string {
	st := arraystack.New()
	ans := ""

	for _, char := range expr {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			ans += string(char)
		} else if char == '(' {
			st.Push(char)
		} else if char == ')' {
			// pop everything until we find the first opening bracket and push it to answer
			for {
				val, ok := st.Peek()

				if !ok || val == '(' {
					break
				}

				ans += string(val.(rune))
				st.Pop()
			}

			st.Pop()
		} else {
			for top, _ := st.Pop(); !st.Empty() && priority(char) <= priority(top.(rune)); top, _ = st.Pop() {
				ans += top.(string)
			}

			st.Push(char)
		}

	}

	for !st.Empty() {
		top, _ := st.Peek()
		ans += string(top.(rune))
		st.Pop()
	}

	return ans
}

func priority(char rune) int {
	priorities := map[rune]int{
		'^': 3,
		'/': 2,
		'*': 2,
		'+': 1,
		'-': 1,
	}

	if priority, ok := priorities[char]; ok {
		return priority
	}

	return -1
}

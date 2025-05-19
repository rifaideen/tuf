package main

func isValid(s string) bool {
	st := Constructor()
	mapping := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		// check if the char is valid opening parenthesis
		if char == '[' || char == '{' || char == '(' {
			st.Push(char)
			continue
		}

		// stack should not be empty, when we expect to pop
		if st.Empty() {
			return false
		}

		// char mapping should match the popped value
		if mapping[char] != st.Pop() {
			return false
		}
	}

	// stack should be empty, by the time we finish all the operations
	return st.Empty()
}

type MyStack struct {
	values []rune
}

func Constructor() MyStack {
	return MyStack{}
}

func (t *MyStack) Push(x rune) {
	t.values = append(t.values, x)
}

func (t *MyStack) Pop() rune {
	val := t.values[len(t.values)-1]

	t.values = t.values[:len(t.values)-1]

	return val
}

func (t *MyStack) Top() rune {
	return t.values[len(t.values)-1]
}

func (t *MyStack) Empty() bool {
	return len(t.values) == 0
}

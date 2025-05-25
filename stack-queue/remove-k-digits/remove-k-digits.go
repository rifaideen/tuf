package main

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	n := len(num)

	if k == n {
		return "0"
	}

	stack := []byte{
		num[0],
	}

	for i := 1; i < len(num); i++ {
		for len(stack) > 0 && k > 0 && (stack[len(stack)-1]-'0') > (num[i]-'0') {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, num[i])
	}

	sb := strings.Builder{}

	for _, v := range stack {
		sb.WriteByte(v)
	}

	return sb.String()
}

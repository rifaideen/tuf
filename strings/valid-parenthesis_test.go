package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeOuterParentheses(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test case 1",
			s:    "",
			want: "",
		},
		{
			name: "test case 2",
			s:    "(()())(())",
			want: "()()()",
		},
		{
			name: "test case 3",
			s:    "(()())(())(()(()))",
			want: "()()()()(())",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeOuterParentheses(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

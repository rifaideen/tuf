package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStack(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    []int
		want []int
	}{
		{
			name: "test case 1",
			s:    []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "single element stack",
			s:    []int{42},
			want: []int{42},
		},
		{
			name: "already reversed stack",
			s:    []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "two elements",
			s:    []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "empty stack",
			s:    []int{},
			want: []int{},
		},
		{
			name: "stack with negative and positive numbers",
			s:    []int{-3, -2, 0, 2, 5},
			want: []int{5, 2, 0, -2, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()

			for _, v := range tt.s {
				s.Push(v)
			}

			ReverseStack(s)

			assert.Equal(t, tt.want, s.items)
		})
	}
}

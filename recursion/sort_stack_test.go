package main

import (
	"reflect"
	"testing"
)

func TestSortStack(t *testing.T) {
	tests := []struct {
		name     string // description of this test case
		input    []int  // values to push onto stack
		expected []int  // expected stack after sorting (bottom to top)
	}{
		{
			name:     "Empty stack",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element stack",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted stack",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted stack",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Mixed order stack",
			input:    []int{3, 1, 5, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Stack with duplicates",
			input:    []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "Stack with negative numbers",
			input:    []int{-3, 1, -5, 2, -1},
			expected: []int{-5, -3, -1, 1, 2},
		},
		{
			name:     "Large stack",
			input:    []int{23, 10, 45, 2, 67, 8, 34, 12, 76, 5},
			expected: []int{2, 5, 8, 10, 12, 23, 34, 45, 67, 76},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create and populate the stack
			s := NewStack()
			for i := len(tt.input) - 1; i >= 0; i-- {
				s.Push(tt.input[i])
			}

			// Sort the stack
			SortStack(s)

			// Check if the stack is sorted correctly
			result := []int{}
			for !s.IsEmpty() {
				val, _ := s.Pop()
				result = append([]int{val}, result...)
			}

			// Compare the result with expected
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortStack() = %v, want %v", result, tt.expected)
			}
		})
	}
}

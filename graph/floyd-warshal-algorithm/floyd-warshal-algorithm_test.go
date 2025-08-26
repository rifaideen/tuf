package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortedDistance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix   *[][]int
		expected [][]int
	}{
		{
			name: "Test case with positive weights and unreachable paths",
			matrix: &[][]int{
				{0, 5, -1, 10},
				{-1, 0, 3, -1},
				{-1, -1, 0, 1},
				{-1, -1, -1, 0},
			},
			expected: [][]int{
				{0, 5, 8, 9},
				{-1, 0, 3, 4},
				{-1, -1, 0, 1},
				{-1, -1, -1, 0},
			},
		},
		{
			name: "A more complex graph with indirect paths",
			matrix: &[][]int{
				{0, 3, 6, -1},
				{3, 0, 2, -1},
				{6, 2, 0, 1},
				{-1, -1, 1, 0},
			},
			expected: [][]int{
				{0, 3, 5, 6},
				{3, 0, 2, 3},
				{5, 2, 0, 1},
				{6, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortedDistance(tt.matrix)
			assert.Equal(t, tt.expected, *tt.matrix)
		})
	}
}

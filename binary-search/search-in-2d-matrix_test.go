package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "test case 1",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
		{
			name: "test case 2",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_searchTarget(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		row    []int
		target int
		want   bool
	}{
		{
			name:   "test case 1",
			row:    []int{1, 2, 3, 4, 5},
			target: 4,
			want:   true,
		},
		{
			name:   "test case 1",
			row:    []int{1, 2, 3, 4, 5},
			target: 6,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchTarget(tt.row, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

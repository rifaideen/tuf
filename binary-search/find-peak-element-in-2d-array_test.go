package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakGrid(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mat  [][]int
		want []int
	}{
		{
			name: "test case 1",
			mat: [][]int{
				{1, 4},
				{3, 2},
			},
			want: []int{0, 1},
		},
		{
			name: "test case 2",
			mat: [][]int{
				{10, 20, 15},
				{21, 30, 14},
				{7, 16, 32},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakGrid(tt.mat)
			assert.NotEqual(t, []int{-1, -1}, got)
		})
	}
}

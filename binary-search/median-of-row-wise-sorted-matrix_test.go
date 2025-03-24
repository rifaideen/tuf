package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_medianOfRowiseSortedMatrix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix [][]int
		want   int
	}{
		{
			name: "test case 1",
			matrix: [][]int{
				{1, 3, 5},
				{2, 6, 9},
				{3, 6, 9},
			},
			want: 5,
		},
		{
			name: "test case 2",
			matrix: [][]int{
				{1},
				{2},
				{3},
			},
			want: 2,
		},
		{
			name: "test case 3",
			matrix: [][]int{
				{3},
				{5},
				{8},
			},
			want: 5,
		},
		{
			name: "test case 4",
			matrix: [][]int{
				{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "test case 5",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{8, 9, 11, 12, 13},
				{21, 23, 25, 27, 29},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := medianOfRowiseSortedMatrix(tt.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}

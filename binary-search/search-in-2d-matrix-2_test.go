package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrixBS2(t *testing.T) {
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
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{10, 12, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
		{
			name: "test case 2",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{10, 12, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrixBS2(tt.matrix, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orangesRotting(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{2, 1, 1, 0, 1, 1},
				{1, 1, 0, 2, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{1, 0, 0, 1, 1, 1},
				{1, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := orangesRotting(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}

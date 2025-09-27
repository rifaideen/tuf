package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxChocolates(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "1x1 Grid",
			grid: [][]int{{5}},
			want: 5,
		},
		{
			name: "2x2 Grid",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 10, // 1+2 + 3+4 = 10
		},
		{
			name: "3x3 All Ones",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 6, // 2 per row × 3 rows (never meet)
		},
		{
			name: "3x3 Peak in Center",
			grid: [][]int{
				{1, 0, 1},
				{0, 10, 0},
				{1, 0, 1},
			},
			want: 14, // 1+1 + 10 + 1+1 = 14
		},
		{
			name: "All Zeros 2x2",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			want: 0,
		},
		{
			name: "3x3 Decreasing",
			grid: [][]int{
				{9, 8, 7},
				{6, 5, 4},
				{3, 2, 1},
			},
			want: 32, // Verified optimal
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChocolates(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}

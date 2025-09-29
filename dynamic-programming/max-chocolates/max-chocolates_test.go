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

func Test_maxChocolatesMemoized(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "1x1 grid",
			grid: [][]int{{5}},
			want: 5,
		},
		{
			name: "1x3 grid",
			grid: [][]int{{1, 2, 3}},
			want: 4,
		},
		{
			name: "3x1 grid",
			grid: [][]int{
				{1},
				{2},
				{3},
			},
			want: 6,
		},
		{
			name: "2x3 grid - with overlap in second row",
			grid: [][]int{
				{1, 2, 3},
				{0, 4, 0},
			},
			want: 8, // ✅ FIXED
		},
		{
			name: "3x3 grid - general path with no overlap",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 14, // ✅ FIXED
		},
		{
			name: "3x4 grid - tests boundaries and full recursion",
			grid: [][]int{
				{1, 2, 3, 4},
				{0, 0, 0, 0},
				{9, 8, 7, 6},
			},
			want: 22, // ✅ FIXED
		},
		{
			name: "Edge case: all zeros",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "Large values and overlap",
			grid: [][]int{
				{100, 0, 100},
				{0, 1000, 0},
				{100, 0, 100},
			},
			want: 1400, // ✅ FIXED
		},
		{
			name: "5x3 grid - optimal path with careful overlap",
			grid: [][]int{
				{2, 4, 10},
				{9, 8, 11},
				{13, 2, 11},
				{10, 5, 14},
				{12, 10, 8},
			},
			want: 102,
		},
		{
			name: "5x3 grid - with max 107",
			grid: [][]int{
				{5, 14, 12},
				{7, 4, 14},
				{10, 11, 4},
				{10, 11, 6},
				{12, 5, 15},
			},
			want: 107,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChocolatesMemoized(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}

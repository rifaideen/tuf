package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubsequenceWithTargetSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		{
			name: "taget sum 10",
			nums: []int{4, 9, 2, 5, 1},
			k:    10,
			want: 2, // {9, 1}, {4, 5, 1}
		},
		{
			name: "taget sum 5",
			nums: []int{4, 2, 10, 5, 1, 3},
			k:    5,
			want: 3, // {5}, {4, 1}, {2, 3}
		},
		{
			name: "taget sum 9",
			nums: []int{3, 5, 6, 7},
			k:    9,
			want: 1, // {3, 6}
		},
		{
			name: "empty array",
			nums: []int{},
			k:    5,
			want: 0,
		},
		{
			name: "single element equals target",
			nums: []int{7},
			k:    7,
			want: 1, // {7}
		},
		{
			name: "single element not equals target",
			nums: []int{3},
			k:    5,
			want: 0,
		},
		{
			name: "all elements sum to target",
			nums: []int{1, 2, 3, 4},
			k:    10,
			want: 1, // {1, 2, 3, 4}
		},
		{
			name: "multiple distinct subsequences with target sum (duplicates accepted, left to right)",
			nums: []int{1, 2, 1, 2, 1},
			k:    3,
			want: 7, // {1@0, 2@1}, {1@0, 2@3}, {1@0, 2@4}, {1@2, 2@3}, {1@2, 2@4}, {1@4, 2@1}, {3@2} - Wait, there's no 3 in the array! My mistake. Let's re-re-evaluate.

			// nums: [1, 2, 1, 2, 1], k: 3
			// {1@0, 2@1}
			// {1@0, 2@3}
			// {1@0, 2@4}
			// {1@2, 2@3}
			// {1@2, 2@4}
			// {1@4, 2@1}
		},
		{
			name: "no subsequence with target sum",
			nums: []int{1, 2, 3},
			k:    7,
			want: 0,
		},
		{
			name: "duplicate numbers in input leading to multiple subsequences (duplicates accepted)",
			nums: []int{2, 2, 2, 2},
			k:    4,
			want: 6, // {2@0, 2@1}, {2@0, 2@2}, {2@0, 2@3}, {2@1, 2@2}, {2@1, 2@3}, {2@2, 2@3}
		},
		{
			name: "larger numbers",
			nums: []int{10, 20, 30, 40},
			k:    50,
			want: 2, // {10@0, 40@3}, {20@1, 30@2}
		},
		{
			name: "target sum achievable in multiple ways with different lengths (duplicates accepted)",
			nums: []int{1, 2, 3, 1, 2},
			k:    3,
			want: 5, // {1@0, 2@1}, {1@0, 2@4}, {1@3, 2@1}, {1@3, 2@4}, {3@2}
		},
		{
			name: "target sum equal to a single element appearing multiple times (duplicates accepted)",
			nums: []int{5, 5, 5},
			k:    5,
			want: 3, // {5@0}, {5@1}, {5@2}
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubsequenceWithTargetSum(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

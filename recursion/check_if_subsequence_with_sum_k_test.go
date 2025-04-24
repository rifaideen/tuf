package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSubsequenceSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want bool
	}{
		{
			name: "tc 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    8,
			want: true, // 3 + 5 = 8
		},
		{
			name: "tc 2",
			nums: []int{4, 3, 9, 2},
			k:    10,
			want: false, // No combination sums to 10
		},
		{
			name: "tc 3",
			nums: []int{10, 20, 15, 5},
			k:    25,
			want: true, // 10 + 15 = 25
		},
		{
			name: "tc 4",
			nums: []int{1, 1, 1, 1, 1},
			k:    3,
			want: true, // 1 + 1 + 1 = 3
		},
		{
			name: "tc 5",
			nums: []int{100},
			k:    100,
			want: true, // Single element matches k
		},
		{
			name: "tc 6",
			nums: []int{2, 4, 6, 8},
			k:    5,
			want: false, // No combination gives 5
		},
		{
			name: "tc 7",
			nums: []int{5, 5, 5, 5},
			k:    10,
			want: true, // 5 + 5 = 10
		},
		{
			name: "tc 8",
			nums: []int{1, 2, 4, 8, 16},
			k:    31,
			want: true, // 1 + 2 + 4 + 8 + 16 = 31 (all elements)
		},
		{
			name: "tc 9",
			nums: []int{1, 2, 4, 8, 16},
			k:    32,
			want: false, // No combination sums to 32
		},
		{
			name: "tc 10",
			nums: []int{200, 300, 500, 1000},
			k:    2000,
			want: true, // 200 + 300 + 500 + 1000 = 2000
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckSubsequenceSum(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

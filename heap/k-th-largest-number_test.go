package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		// {
		// 	name: "tc 1",
		// 	nums: []int{3, 2, 1, 5, 6, 4},
		// 	k:    2,
		// 	want: 5,
		// },
		{
			name: "tc 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

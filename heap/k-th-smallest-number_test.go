package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSmallest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		{
			name: "tc 1",
			nums: []int{6, 5, 4, 3, 2, 1},
			k:    2,
			want: 2,
		},
		{
			name: "tc 2",
			nums: []int{1, 2, 3, 8, 7, 9},
			k:    5,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSmallest(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

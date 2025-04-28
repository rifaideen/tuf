package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "test case 1",
			nums: []int{5, 2, 1},
			want: []int{0, 1, 2, 3, 5, 6, 7, 8},
		},
		{
			name: "test case 2",
			nums: []int{3, 1, 2},
			want: []int{0, 1, 2, 3, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetSum(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

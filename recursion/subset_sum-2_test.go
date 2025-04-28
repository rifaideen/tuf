package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetSumTwo(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want [][]int
	}{
		{
			name: "test case 1",
			nums: []int{1, 2, 2},
			want: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetSumTwo(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

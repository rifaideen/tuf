package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replaceByRank(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			nums: []int{20, 15, 26, 2, 98, 6},
			want: []int{4, 3, 5, 1, 6, 2},
		},
		{
			nums: []int{2, 2, 1, 6},
			want: []int{2, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceByRank(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}

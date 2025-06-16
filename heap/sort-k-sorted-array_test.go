package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nearlySorted(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
		k    int
	}{
		{
			nums: []int{6, 5, 3, 2, 8, 10, 9},
			k:    3,
			want: []int{2, 3, 5, 6, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nearlySorted(tt.nums, tt.k)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}

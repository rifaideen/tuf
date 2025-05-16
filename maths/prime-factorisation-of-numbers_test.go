package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_primeFactors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want [][]int
	}{
		{
			name: "test case 1",
			nums: []int{2, 3, 4, 5, 6},
			want: [][]int{
				{2},
				{3},
				{2, 2},
				{5},
				{2, 3},
			},
		},
		{
			name: "test case 2",
			nums: []int{7, 12, 18},
			want: [][]int{
				{7},
				{2, 2, 3},
				{2, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := primeFactors(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

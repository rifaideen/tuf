package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		k    int
		n    int
		want [][]int
	}{
		{
			name: "combination sum 7 with length 3",
			k:    3,
			n:    7,
			want: [][]int{
				{1, 2, 4},
			},
		},
		{
			name: "combination sum 9 with length 3",
			k:    3,
			n:    9,
			want: [][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
		{
			name: "combination sum 1 with length 4",
			k:    4,
			n:    1,
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

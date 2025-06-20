package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxCombinations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		N    int
		K    int
		A    []int
		B    []int
		Want []int
	}{
		{
			name: "tc 1",
			N:    2,
			K:    2,
			A:    []int{3, 2},
			B:    []int{1, 4},
			Want: []int{7, 6},
		},
		{
			name: "tc 2",
			N:    4,
			K:    3,
			A:    []int{1, 4, 2, 3},
			B:    []int{2, 5, 1, 6},
			Want: []int{10, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCombinations(tt.N, tt.K, tt.A, tt.B)
			assert.Equal(t, tt.Want, got)
		})
	}
}

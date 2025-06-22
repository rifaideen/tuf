package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ropes []int
		want  int
	}{
		{
			name:  "tc 1",
			ropes: []int{4, 3, 2, 6},
			want:  29,
		},
		{
			name:  "tc 2",
			ropes: []int{4, 2, 7, 6, 9},
			want:  62,
		},
		{
			name:  "tc 3",
			ropes: []int{10},
			want:  0,
		},
		{
			name:  "tc 4",
			ropes: []int{5, 10, 2, 1, 6},
			want:  49,
		},
		{
			name:  "tc 5",
			ropes: []int{100, 500, 50, 25, 8, 6, 3, 0},
			want:  1047,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost(tt.ropes)
			assert.Equal(t, tt.want, got)
		})
	}
}

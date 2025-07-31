package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bfsTraversal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n     int
		start int
		list  [][]int
		want  []int
	}{
		{
			n:     10,
			start: 1,
			list: [][]int{
				{1, 2},
				{1, 7},
				{2, 3},
				{2, 6},
				{7, 8},
				{7, 10},
				{3, 4},
				{3, 5},
				{8, 9},
			},
			want: []int{1, 2, 7, 3, 6, 8, 10, 4, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bfsTraversal(tt.n, tt.start, tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

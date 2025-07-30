package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dfsTraversal(t *testing.T) {
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
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			n:     5,
			start: 1,
			list: [][]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{2, 5},
			},
			want: []int{1, 2, 4, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dfsTraversal(tt.n, tt.start, tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

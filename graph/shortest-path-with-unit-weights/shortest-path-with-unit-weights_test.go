package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		edges [][]int
		n     int
		m     int
		src   int
		want  []int
	}{
		{
			edges: [][]int{
				{0, 1},
				{0, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{1, 2},
				{2, 6},
				{6, 7},
				{7, 8},
				{6, 8},
			},
			n:    9,
			m:    10,
			want: []int{0, 1, 2, 1, 2, 3, 3, 4, 4},
			src:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestPath(tt.edges, tt.n, tt.m, tt.src)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n     int
		m     int
		edges [][]int
		want  []int
	}{
		{
			n: 6,
			m: 7,
			edges: [][]int{
				{0, 1, 2},
				{0, 4, 1},
				{4, 5, 4},
				{4, 2, 2},
				{1, 2, 3},
				{2, 3, 6},
				{5, 3, 1},
			},
			want: []int{0, 2, 3, 6, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestPath(tt.n, tt.m, tt.edges)
			assert.Equal(t, tt.want, got)
		})
	}
}

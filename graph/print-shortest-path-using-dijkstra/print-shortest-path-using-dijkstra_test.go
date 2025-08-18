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
			n: 5,
			m: 6,
			edges: [][]int{
				{1, 2, 2},
				{2, 5, 5},
				{2, 3, 4},
				{1, 4, 1},
				{4, 3, 3},
				{3, 5, 1},
			},
			want: []int{1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestPath(tt.n, tt.m, tt.edges)
			assert.Equal(t, tt.want, got)
		})
	}
}

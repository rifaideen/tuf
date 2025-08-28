package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spanningTree(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		v     int
		edges [][]int
		want  int
	}{
		{
			v: 4,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 2},
				{2, 3, 3},
				{0, 3, 4},
			},
			want: 6,
		},
		{
			v: 4,
			edges: [][]int{
				{0, 1, 4},
				{0, 2, 3},
				{1, 2, 1},
				{1, 3, 2},
				{2, 3, 5},
			},
			want: 6,
		},
		{
			v: 5,
			edges: [][]int{
				{0, 1, 2},
				{0, 3, 6},
				{1, 2, 3},
				{1, 3, 8},
				{1, 4, 5},
				{2, 4, 7},
			},
			want: 16,
		},
		{
			v:     1,
			edges: [][]int{},
			want:  0,
		},
		{
			v: 4,
			edges: [][]int{
				{0, 1, 10},
				{0, 2, 6},
				{0, 3, 5},
				{1, 3, 15},
				{2, 3, 4},
			},
			want: 19,
		},
		{
			v: 5,
			edges: [][]int{
				{0, 1, 1},
				{0, 2, 1},
				{0, 3, 1},
				{0, 4, 1},
			},
			want: 4,
		},
		{
			v: 6,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 2},
				{2, 3, 3},
				{3, 4, 4},
				{4, 5, 5},
				{0, 5, 15},
			},
			want: 15,
		},
		{
			v: 4,
			edges: [][]int{
				{0, 1, 7},
				{0, 2, 6},
				{0, 3, 5},
				{1, 2, 4},
				{2, 3, 3},
				{1, 3, 2},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spanningTree(tt.v, tt.edges)
			assert.Equal(t, tt.want, got)
		})
	}
}

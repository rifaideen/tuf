package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_articulationPoints(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n       int
		adjList [][]int
		want    int
	}{
		{
			n: 7,
			adjList: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{2, 4},
				{2, 5},
				{3, 2},
				{4, 6},
				{5, 6},
			},
			want: 2,
		},
		{
			n: 5,
			adjList: [][]int{
				{0, 1},
				{1, 4},
				{2, 4},
				{3, 4},
				{2, 3},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := articulationPoints(tt.n, tt.adjList)
			assert.Equal(t, tt.want, got)
		})
	}
}

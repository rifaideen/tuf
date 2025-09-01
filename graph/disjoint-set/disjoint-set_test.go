package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjointSet_UnionByRank(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		n      int
		unions [][]int
	}{
		{
			n: 7,
			unions: [][]int{
				{1, 2},
				{2, 3},
				{4, 5},
				{6, 7},
				{5, 6},
				{3, 7, 0}, // check if same nodes belongs to same parent, expected false
				{3, 7},
				{3, 7, 1}, // check if same nodes belongs to same parent, expected true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := New(tt.n)

			for _, node := range tt.unions {
				u := node[0]
				v := node[1]

				if len(node) == 2 {
					ds.UnionByRank(u, v)
				} else {
					got := ds.FindUltimateParent(u) == ds.FindUltimateParent(v)

					assert.Equal(t, node[2] == 1, got)
				}
			}
		})
	}
}

func TestDisjointSet_UnionBySize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		n      int
		unions [][]int
	}{
		{
			n: 7,
			unions: [][]int{
				{1, 2},
				{2, 3},
				{4, 5},
				{6, 7},
				{5, 6},
				{3, 7, 0}, // check if same nodes belongs to same parent, expected false
				{3, 7},
				{3, 7, 1}, // check if same nodes belongs to same parent, expected true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := New(tt.n)

			for _, node := range tt.unions {
				u := node[0]
				v := node[1]

				if len(node) == 2 {
					ds.UnionBySize(u, v)
				} else {
					got := ds.FindUltimateParent(u) == ds.FindUltimateParent(v)

					assert.Equal(t, node[2] == 1, got)
				}
			}
		})
	}
}

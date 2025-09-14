package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kosarajuAlgorithm(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		v       int
		adjList [][]int
		want    int
	}{
		{
			name: "Simple Connected Graph",
			v:    4,
			adjList: [][]int{
				{1},
				{2},
				{3},
				{0},
			},
			want: 1, // All nodes form a single cycle → one SCC
		},
		{
			name: "Graph with Two Components",
			v:    5,
			adjList: [][]int{
				{1},
				{2},
				{0},
				{4},
				{3},
			},
			want: 2, // SCC1: {0,1,2} (cycle), SCC2: {3,4} (cycle)
		},
		{
			name: "Disconnected Graph with No Cycles",
			v:    6,
			adjList: [][]int{
				{1}, // 0 → 1
				{},  // 1 → none
				{3}, // 2 → 3
				{2}, // 3 → 2
				{5}, // 4 → 5
				{},  // 5 → none
			},
			want: 5, // Five SCCs: {0}, {1}, {2,3}, {4}, {5}
		},
		{
			name: "Single Node",
			v:    1,
			adjList: [][]int{
				{},
			},
			want: 1, // Single node is its own SCC
		},
		{
			name: "No Edges",
			v:    3,
			adjList: [][]int{
				{},
				{},
				{},
			},
			want: 3, // Each node is its own SCC
		},
		{
			name: "Complex with Nested SCCs",
			v:    7,
			adjList: [][]int{
				{1},
				{2},
				{1, 3}, // 2 → 1 (cycle) and 2 → 3
				{4},
				{5},
				{6},
				{3}, // 6 → 3; so 3→4→5→6→3 → cycle
			},
			want: 3, // SCC1: {1,2}, SCC2: {3,4,5,6}, SCC3: {0}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kosarajuAlgorithm(tt.v, tt.adjList)
			assert.Equal(t, tt.want, got)
		})
	}
}

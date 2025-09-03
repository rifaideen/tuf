package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spanningTree(t *testing.T) {
	tests := []struct {
		name    string    // description of this test case
		v       int       // number of vertices
		adjList [][][]int // adjacency list: adjList[u] = [[v, w], ...]
		want    int       // expected MST total weight
	}{
		{
			name: "single node - no edges",
			v:    1,
			adjList: [][][]int{
				{}, // node 0 has no neighbors
			},
			want: 0, // MST of one node has zero weight
		},
		{
			name: "two nodes - one edge",
			v:    2,
			adjList: [][][]int{
				{{1, 5}}, // 0 --5-- 1
				{{0, 5}}, // undirected
			},
			want: 5, // Only one edge → must include it
		},
		{
			name: "triangle - choose two smallest edges",
			v:    3,
			adjList: [][][]int{
				{{1, 1}, {2, 4}}, // 0--1 (1), 0--2 (4)
				{{0, 1}, {2, 2}}, // 1--0 (1), 1--2 (2)
				{{0, 4}, {1, 2}}, // 2--0 (4), 2--1 (2)
			},
			want: 3, // MST: 0--1 (1) + 1--2 (2) = 3. Skip 0--2 (4)
		},
		{
			name: "linear chain (path graph)",
			v:    4,
			adjList: [][][]int{
				{{1, 2}},
				{{0, 2}, {2, 3}},
				{{1, 3}, {3, 1}},
				{{2, 1}},
			},
			// Edges: 0--1(2), 1--2(3), 2--3(1)
			// MST must include all (it's already a tree)
			want: 2 + 3 + 1, // = 6
		},
		{
			name: "square with diagonal - choose lightest three edges",
			v:    4,
			adjList: [][][]int{
				{{1, 1}, {3, 4}},         // 0--1(1), 0--3(4)
				{{0, 1}, {2, 2}, {3, 3}}, // 1--0(1), 1--2(2), 1--3(3)
				{{1, 2}, {3, 1}},         // 2--1(2), 2--3(1)
				{{0, 4}, {1, 3}, {2, 1}}, // 3--0(4), 3--1(3), 3--2(1)
			},
			// Best MST: 0--1(1), 1--2(2), 2--3(1) → total = 4
			// Or: 0--1(1), 1--2(2), 3--2(1) → same
			want: 4,
		},
		{
			name: "star graph - center connected to all",
			v:    4,
			adjList: [][][]int{
				{{1, 5}, {2, 3}, {3, 7}}, // center node 0
				{{0, 5}},
				{{0, 3}},
				{{0, 7}},
			},
			// MST is the graph itself (already a tree)
			// But we can't skip any → total = 5+3+7 = 15
			want: 15,
		},
		{
			name: "larger graph (classic MST example)",
			v:    5,
			adjList: [][][]int{
				{{1, 2}, {3, 6}},                 // 0
				{{0, 2}, {2, 3}, {3, 8}, {4, 5}}, // 1
				{{1, 3}, {4, 7}},                 // 2
				{{0, 6}, {1, 8}, {4, 9}},         // 3
				{{1, 5}, {2, 7}, {3, 9}},         // 4
			},
			// Expected MST edges:
			// 0--1 (2), 1--2 (3), 1--4 (5), 0--3 (6) → total = 2+3+5+6 = 16
			// (Alternative: 4--1 instead of 4--2, etc. — weight same)
			want: 16,
		},
		{
			name: "duplicate edges with different weights (should pick lighter)",
			v:    3,
			adjList: [][][]int{
				{{1, 5}, {1, 2}, {2, 4}}, // 0 has two edges to 1: 5 and 2 → pick 2
				{{0, 2}, {0, 5}, {2, 1}}, // 1--0 (2), and 1--2 (1)
				{{0, 4}, {1, 1}},         // 2--0 (4), 2--1 (1)
			},
			// Lightest valid: 0--1 (2), 1--2 (1) → total = 3
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spanningTree(tt.v, tt.adjList)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_graphColoring(t *testing.T) {
	// Test case 1: A cycle of 4 nodes (can be colored with 2 colors)
	graph1 := [101][101]int{}
	graph1[0][1] = 1
	graph1[1][0] = 1
	graph1[1][2] = 1
	graph1[2][1] = 1
	graph1[2][3] = 1
	graph1[3][2] = 1
	graph1[3][0] = 1
	graph1[0][3] = 1

	// Test case 2: A complete graph of 3 nodes (needs 3 colors)
	graph2 := [101][101]int{}
	graph2[0][1] = 1
	graph2[1][0] = 1
	graph2[0][2] = 1
	graph2[2][0] = 1
	graph2[1][2] = 1
	graph2[2][1] = 1

	// Test case 3: A disconnected graph (two separate edges)
	graph3 := [101][101]int{}
	graph3[0][1] = 1
	graph3[1][0] = 1
	graph3[2][3] = 1
	graph3[3][2] = 1

	// Test case 4: A single node (can be colored with 1 color)
	graph4 := [101][101]int{}

	// Test case 5: A line of 3 nodes (can be colored with 2 colors)
	graph5 := [101][101]int{}
	graph5[0][1] = 1
	graph5[1][0] = 1
	graph5[1][2] = 1
	graph5[2][1] = 1

	tests := []struct {
		name  string
		graph [101][101]int
		m     int
		n     int
		want  bool
	}{
		{
			name:  "Cycle of 4 - 3 colors",
			graph: graph1,
			m:     3,
			n:     4,
			want:  true,
		},
		{
			name:  "Cycle of 4 - 2 colors",
			graph: graph1,
			m:     2,
			n:     4,
			want:  true,
		},
		{
			name:  "Cycle of 4 - 1 color",
			graph: graph1,
			m:     1,
			n:     4,
			want:  false,
		},
		{
			name:  "Complete graph of 3 - 3 colors",
			graph: graph2,
			m:     3,
			n:     3,
			want:  true,
		},
		{
			name:  "Complete graph of 3 - 2 colors",
			graph: graph2,
			m:     2,
			n:     3,
			want:  false,
		},
		{
			name:  "Disconnected graph - 2 colors",
			graph: graph3,
			m:     2,
			n:     4,
			want:  true,
		},
		{
			name:  "Disconnected graph - 1 color",
			graph: graph3,
			m:     1,
			n:     4,
			want:  false,
		},
		{
			name:  "Single node - 1 color",
			graph: graph4,
			m:     1,
			n:     1,
			want:  true,
		},
		{
			name:  "Single node - 0 colors",
			graph: graph4,
			m:     0,
			n:     1,
			want:  false,
		},
		{
			name:  "Line of 3 - 2 colors",
			graph: graph5,
			m:     2,
			n:     3,
			want:  true,
		},
		{
			name:  "Line of 3 - 1 color",
			graph: graph5,
			m:     1,
			n:     3,
			want:  false,
		},
		{
			name:  "Empty graph - 2 colors",
			graph: [101][101]int{},
			m:     2,
			n:     0,
			want:  true, // Technically true, no nodes to color
		},
		{
			name:  "Empty graph - 0 colors",
			graph: [101][101]int{},
			m:     0,
			n:     0,
			want:  true, // Technically true, no nodes to color
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphColoring(tt.graph, tt.m, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

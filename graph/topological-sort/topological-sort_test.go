package main

import (
	"fmt"
	"testing"
)

func Test_topoSort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		vertices int
		adjList  [][]int
		want     []int
	}{
		{
			vertices: 4,
			adjList: [][]int{
				{0, 0},
				{3, 0},
				{1, 0},
				{2, 0},
			},
		},
		{
			vertices: 6,
			adjList: [][]int{
				{1, 3},
				{2, 3},
				{4, 1},
				{4, 0},
				{5, 0},
				{5, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topoSort(tt.vertices, tt.adjList)
			fmt.Println(got)
		})
	}
}

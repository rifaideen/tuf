package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]string
		want int
	}{
		{
			grid: [][]string{
				{"L", "L", "W", "W", "W"},
				{"W", "L", "W", "W", "L"},
				{"L", "W", "W", "L", "L"},
				{"W", "W", "W", "W", "W"},
				{"L", "W", "L", "L", "W"},
			},
			want: 4,
		},
		{
			grid: [][]string{
				{"W", "L", "L", "L", "W", "W", "W"},
				{"W", "W", "L", "L", "W", "L", "W"},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}

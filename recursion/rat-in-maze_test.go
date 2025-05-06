package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rarInMaze(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board [][]int
		want  []string
	}{
		{
			name: "test case 1",
			board: [][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 1},
				{1, 1, 0, 0},
				{0, 1, 1, 1},
			},
			want: []string{
				"DDRDRR",
				"DRDDRR",
			},
		},
		{
			name: "test case 2",
			board: [][]int{
				{1, 0},
				{1, 0},
			},
			want: []string{},
		},
		{
			name: "test case 3",
			board: [][]int{
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{1, 1, 1, 1},
			},
			want: []string{
				"DDDRRR",
				"DRRRDD",
				"DRRURDDD",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ratInMaze(tt.board)
			assert.Equal(t, tt.want, got)
		})
	}
}

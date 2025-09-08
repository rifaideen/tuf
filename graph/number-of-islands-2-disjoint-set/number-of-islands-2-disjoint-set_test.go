package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name      string // description of this test case
		n         int    // number of rows
		m         int    // number of columns
		operators [][]int
		want      []int
	}{
		{
			name: "Single island formed step by step in a 3x3 grid",
			n:    3, // rows
			m:    3, // cols
			operators: [][]int{
				{0, 0}, // add (0,0) → 1
				{0, 1}, // connects to (0,0) → merge → 1
				{1, 1}, // connects to (0,1) → merge → 1
				{1, 0}, // connects to both → merge → 1
			},
			want: []int{1, 1, 1, 1}, // all part of one island
		},
		{
			name: "Two separate islands grow independently in 2x2 grid",
			n:    2,
			m:    2,
			operators: [][]int{
				{0, 0}, // → 1
				{1, 1}, // not adjacent to (0,0) → 2
				{0, 1}, // now connects both → merge → 1
			},
			want: []int{1, 2, 1},
		},
		{
			name: "Duplicates and isolated lands in 3x3 grid",
			n:    3,
			m:    3,
			operators: [][]int{
				{0, 0}, // → 1
				{0, 0}, // duplicate → still 1
				{2, 2}, // far away → 2
				{1, 1}, // middle → not touching yet → 3
				{1, 1}, // duplicate → still 3
				{0, 1}, // connects (0,0) and (1,1)? Not yet → (0,1) touches (0,0), but (1,1) not adjacent → now (0,1) connects (0,0) and (1,1)? Wait:
				// (0,1) is adjacent to (0,0) and (1,1)? No → (1,1) is down-right → diagonal.
				// So (0,1) connects to (0,0), but not (1,1) → islands: [0,0+0,1], [1,1], [2,2] → still 3
				{1, 0}, // connects (0,0) and (1,1)? (1,0) is adjacent to (0,0) and (1,1)? Yes: down from (0,0), left from (1,1)? No → (1,0) adjacent to (0,0) (up) and (1,1) (right)? Yes → so now (1,0) connects them!
				// So when we add (1,0), it connects to (0,0) (already in same comp), and to (1,1) → merge → counter drops from 3 to 2
				// But wait — (1,0) connects (0,0) and (1,1)? Let's see:
				// - (1,0) up → (0,0): land → union
				// - (1,0) right → (1,1): land → union
				// So (1,0) merges the two components: [0,0], [1,1] → now one component
				// So total islands: [merged], [2,2] → 2
			},
			want: []int{1, 1, 2, 3, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.n, tt.m, tt.operators)
			assert.Equal(t, tt.want, got)
		})
	}
}

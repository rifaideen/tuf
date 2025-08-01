package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		v    int
		e    int
		list [][]int
		want bool
	}{
		{
			v: 8,
			e: 7,
			list: [][]int{
				0: {},
				1: {2, 3},
				2: {1, 5, 6},
				3: {1, 4, 7},
				4: {3},
				5: {2},
				6: {2},
				7: {3, 8},
				8: {7},
			},
			want: false,
		},
		{
			v: 8,
			e: 6,
			list: [][]int{
				0: {},
				1: {2},
				2: {1, 3},
				3: {2},
				4: {5, 6},
				5: {4, 6},
				6: {4, 5},
				7: {8},
				8: {7},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.v, tt.e, tt.list)
			assert.Equal(t, tt.want, got)

			got2 := detectCycleDFS(tt.v, tt.e, tt.list)
			assert.Equal(t, tt.want, got2)
		})
	}
}

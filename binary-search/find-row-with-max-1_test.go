package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countOnes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		row  []int
		want int
	}{
		{
			name: "test case 1",
			row:  []int{0, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "test case 2",
			row:  []int{1, 1, 1, 1, 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countOnes(tt.row)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaxOnesRow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix [][]int
		want   int
	}{
		{
			name: "test case 1",
			matrix: [][]int{
				{
					1, 1, 1,
				},
				{
					0, 0, 1,
				},
				{
					0, 0, 0,
				},
			},
			want: 0,
		},
		{
			name: "test case 2",
			matrix: [][]int{
				{
					1, 1,
				},
				{
					1, 1,
				},
				{
					0, 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxOnesRow(tt.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}

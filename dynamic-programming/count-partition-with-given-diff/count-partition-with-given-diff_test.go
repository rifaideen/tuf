package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPartitions(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		diff int
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 1, 2, 3},
			diff: 1,
			want: 3,
			n:    4,
		},
		{
			arr:  []int{1, 2, 3, 4},
			diff: 2,
			want: 2,
			n:    4,
		},
		{
			arr:  []int{0, 0, 0},
			diff: 0,
			want: 8,
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitions(tt.n, tt.diff, tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_countPartitionsT(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		diff int
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 1, 2, 3},
			diff: 1,
			want: 3,
			n:    4,
		},
		{
			arr:  []int{1, 2, 3, 4},
			diff: 2,
			want: 2,
			n:    4,
		},
		{
			arr:  []int{0, 0, 0},
			diff: 0,
			want: 8,
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitionsT(tt.n, tt.diff, tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_countPartitionsOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		diff int
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 1, 2, 3},
			diff: 1,
			want: 3,
			n:    4,
		},
		{
			arr:  []int{1, 2, 3, 4},
			diff: 2,
			want: 2,
			n:    4,
		},
		{
			arr:  []int{0, 0, 0},
			diff: 0,
			want: 8,
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitionsOptimized(tt.n, tt.diff, tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

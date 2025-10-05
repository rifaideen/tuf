package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_perfectSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		K    int
		want int
	}{
		{
			arr:  []int{2, 3, 5, 16, 8, 10},
			K:    10,
			want: 3,
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			K:    5,
			want: 3,
		},
		{
			arr:  []int{2, 2, 2, 2},
			K:    4,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := perfectSum(tt.arr, tt.K)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_perfectSumT(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		K    int
		want int
	}{
		{
			arr:  []int{2, 3, 5, 16, 8, 10},
			K:    10,
			want: 3,
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			K:    5,
			want: 3,
		},
		{
			arr:  []int{2, 2, 2, 2},
			K:    4,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := perfectSumT(tt.arr, tt.K)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_perfectSumOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		K    int
		want int
	}{
		{
			arr:  []int{2, 3, 5, 16, 8, 10},
			K:    10,
			want: 3,
		},
		{
			arr:  []int{1, 2, 3, 4, 5},
			K:    5,
			want: 3,
		},
		{
			arr:  []int{2, 2, 2, 2},
			K:    4,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := perfectSumOptimized(tt.arr, tt.K)
			assert.Equal(t, tt.want, got)
		})
	}
}

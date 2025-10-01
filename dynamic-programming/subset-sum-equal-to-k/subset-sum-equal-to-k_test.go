package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsetSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr    []int
		target int
		want   bool
	}{
		{
			arr:    []int{1, 2, 7, 3},
			target: 6,
			want:   true,
		},
		{
			arr:    []int{2, 3, 5},
			target: 6,
			want:   false,
		},
		{
			arr:    []int{7, 54, 4, 12, 15, 5},
			target: 9,
			want:   true,
		},
		{
			name:   "large array with repeated elements - tests memoization under load",
			arr:    []int{2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2, 3, 1, 3, 2, 1, 2, 3, 1, 2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2},
			target: 20,
			want:   true, // achievable via combination of many 1s, 2s, and 3s
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsetSum(tt.arr, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isSubsetSumT(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr    []int
		target int
		want   bool
	}{
		{
			arr:    []int{1, 2, 7, 3},
			target: 6,
			want:   true,
		},
		{
			arr:    []int{2, 3, 5},
			target: 6,
			want:   false,
		},
		{
			arr:    []int{7, 54, 4, 12, 15, 5},
			target: 9,
			want:   true,
		},
		{
			name:   "large array with repeated elements - tests memoization under load",
			arr:    []int{2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2, 3, 1, 3, 2, 1, 2, 3, 1, 2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2},
			target: 20,
			want:   true, // achievable via combination of many 1s, 2s, and 3s
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsetSumT(tt.arr, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isSubsetSumOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr    []int
		target int
		want   bool
	}{
		{
			arr:    []int{1, 2, 7, 3},
			target: 6,
			want:   true,
		},
		{
			arr:    []int{2, 3, 5},
			target: 6,
			want:   false,
		},
		{
			arr:    []int{7, 54, 4, 12, 15, 5},
			target: 9,
			want:   true,
		},
		{
			name:   "large array with repeated elements - tests memoization under load",
			arr:    []int{2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2, 3, 1, 3, 2, 1, 2, 3, 1, 2, 1, 3, 2, 1, 3, 1, 2, 3, 1, 2},
			target: 20,
			want:   true, // achievable via combination of many 1s, 2s, and 3s
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsetSumOptimized(tt.arr, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

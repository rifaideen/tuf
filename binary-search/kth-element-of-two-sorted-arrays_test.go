package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		k     int
		want  int
	}{
		{
			name:  "test case 1",
			nums1: []int{2, 3, 45},
			nums2: []int{4, 6, 7, 8},
			k:     4,
			want:  6,
		},
		{
			name:  "test case 2",
			nums1: []int{1, 2, 3, 5, 6},
			nums2: []int{4, 7, 8, 9, 100},
			k:     6,
			want:  6,
		},
		{
			name:  "test case 3",
			nums1: []int{76, 94, 100},
			nums2: []int{5, 5, 9, 11, 31, 36, 68, 71, 75, 100},
			k:     12,
			want:  100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthElement(tt.nums1, tt.nums2, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKthElementBS(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		k     int
		want  int
	}{
		{
			name:  "test case 1",
			nums1: []int{2, 3, 45},
			nums2: []int{4, 6, 7, 8},
			k:     4,
			want:  6,
		},
		{
			name:  "test case 2",
			nums1: []int{1, 2, 3, 5, 6},
			nums2: []int{4, 7, 8, 9, 100},
			k:     6,
			want:  6,
		},
		{
			name:  "test case 3",
			nums1: []int{76, 94, 100},
			nums2: []int{5, 5, 9, 11, 31, 36, 68, 71, 75, 100},
			k:     12,
			want:  100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthElementBS(tt.nums1, tt.nums2, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

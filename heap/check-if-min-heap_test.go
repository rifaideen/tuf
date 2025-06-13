package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMinHeap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values []int
		want   bool
	}{
		{
			name:   "valid min-heap",
			values: []int{10, 20, 30, 21, 23},
			want:   true,
		},
		{
			name:   "invalid min-heap",
			values: []int{10, 20, 30, 25, 15},
			want:   false,
		},
		{
			name:   "invalid min-heap",
			values: []int{10, 20, 30},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMinHeap(tt.values)
			assert.Equal(t, tt.want, got)
		})
	}
}

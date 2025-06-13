package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMaxheap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values []int
		want   bool
	}{
		{
			name:   "invalid max-heap",
			values: []int{10, 20, 30, 21, 23},
			want:   false,
		},
		{
			name:   "valid max-heap",
			values: []int{90, 15, 10, 7, 12, 2},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMaxheap(tt.values)
			assert.Equal(t, tt.want, got)
		})
	}
}

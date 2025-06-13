package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minToMaxHeap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values []int
		want   []int
	}{
		{
			name:   "test case 1",
			values: []int{10, 20, 30, 21, 23},
			want:   []int{30, 21, 23, 10, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minToMaxHeap(tt.values)
			assert.Equal(t, tt.want, got)
		})
	}
}

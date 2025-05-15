package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRange(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		left  int
		right int
		want  int
	}{
		{
			name:  "test case 1",
			left:  3,
			right: 5,
			want:  2,
		},
		{
			name:  "test case 2",
			left:  1,
			right: 3,
			want:  0,
		},
		{
			name:  "test case 3",
			left:  4,
			right: 7,
			want:  0,
		},
		{
			name:  "test case 3",
			left:  5,
			right: 8,
			want:  12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRange(tt.left, tt.right)
			assert.Equal(t, tt.want, got)
		})
	}
}

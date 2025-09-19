package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_frogJump(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		heights []int
		want    int
	}{
		{
			heights: []int{2, 1, 3, 5, 4},
			want:    2,
		},
		{
			heights: []int{7, 5, 1, 2, 6},
			want:    9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frogJump(tt.heights)
			assert.Equal(t, tt.want, got)

			got = frogJump2(tt.heights)
			assert.Equal(t, tt.want, got)

			got = frogJump3(tt.heights)
			assert.Equal(t, tt.want, got)
		})
	}
}

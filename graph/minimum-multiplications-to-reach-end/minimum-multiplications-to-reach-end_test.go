package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumMultiplications(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr   []int
		start int
		end   int
		want  int
	}{
		{
			arr:   []int{2, 5, 7},
			start: 3,
			end:   30,
			want:  2,
		},
		{
			arr:   []int{3, 4, 65},
			start: 7,
			end:   66175,
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumMultiplications(tt.arr, tt.start, tt.end)
			assert.Equal(t, tt.want, got)
		})
	}
}

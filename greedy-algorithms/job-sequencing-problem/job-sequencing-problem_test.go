package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jobSequencing(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		deadlines []int
		profits   []int
		want      []int
	}{
		{
			deadlines: []int{4, 1, 1, 1},
			profits:   []int{20, 10, 40, 30},
			want:      []int{2, 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jobSequencing(tt.deadlines, tt.profits)
			assert.Equal(t, tt.want, got)
		})
	}
}

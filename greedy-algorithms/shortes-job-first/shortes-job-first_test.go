package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_averageTimes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		jobs []int
		want int
	}{
		{
			jobs: []int{4, 3, 7, 1, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageTimes(tt.jobs)
			assert.Equal(t, tt.want, got)
		})
	}
}

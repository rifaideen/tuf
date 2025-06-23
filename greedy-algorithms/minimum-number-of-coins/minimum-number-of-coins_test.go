package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPartition(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want []int
	}{

		{
			n:    43,
			want: []int{20, 20, 2, 1},
		},
		{
			n:    486,
			want: []int{200, 200, 50, 20, 10, 5, 1},
		},
		{
			n:    1019,
			want: []int{500, 500, 10, 5, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPartition(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

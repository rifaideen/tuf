package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divisors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want []int
	}{
		{
			name: "test case 1",
			n:    6,
			want: []int{1, 6, 2, 3},
		},
		{
			name: "test case 2",
			n:    8,
			want: []int{1, 8, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divisors(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

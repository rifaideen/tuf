package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printAllPrimeFactors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want []int
	}{
		{
			name: "test case 1",
			n:    37,
			want: []int{37},
		},
		{
			name: "test case 1",
			n:    15,
			want: []int{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printAllPrimeFactors(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

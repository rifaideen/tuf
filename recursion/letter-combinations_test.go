package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		digits string
		want   []string
	}{
		{
			name:   "test case 1",
			digits: "23",
			want: []string{
				"ad", "ae", "af",
				"bd", "be", "bf",
				"cd", "ce", "cf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			assert.Equal(t, tt.want, got)
		})
	}
}

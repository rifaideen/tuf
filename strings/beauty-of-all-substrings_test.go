package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_beautySum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := beautySum(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

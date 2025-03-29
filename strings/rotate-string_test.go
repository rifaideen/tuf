package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		goal string
		want bool
	}{
		{
			name: "test case 1",
			s:    "abcde",
			goal: "cdeab",
			want: true,
		},
		{
			name: "test case 2",
			s:    "abcde",
			goal: "abced",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateString(tt.s, tt.goal)
			assert.Equal(t, tt.want, got)
		})
	}
}

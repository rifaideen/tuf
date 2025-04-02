package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "test case 1",
			s:    "42",
			want: 42,
		},
		{
			name: "test case 2",
			s:    "   -042",
			want: -42,
		},
		{
			name: "test case 3",
			s:    "1337c0d3",
			want: 1337,
		},
		{
			name: "test case 4",
			s:    "0-1",
			want: 0,
		},
		{
			name: "test case 5",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "test case 6",
			s:    "words and 987",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

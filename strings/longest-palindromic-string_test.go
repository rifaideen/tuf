package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test case 0",
			s:    "abb",
			want: "bb",
		},
		{
			name: "test case 1",
			s:    "babad",
			want: "aba",
		},
		{
			name: "test case 2",
			s:    "cbbd",
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

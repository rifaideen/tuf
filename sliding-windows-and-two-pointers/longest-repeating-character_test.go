package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		k    int
		want int
	}{
		{
			s:    "AABA",
			want: 2,
			k:    0,
		},
		{
			s:    "AABABBA",
			want: 4,
			k:    1,
		},
		{
			s:    "AABABBA",
			want: 4,
			k:    1,
		},
		{
			s:    "ABAA",
			want: 2,
			k:    0,
		},
		{
			s:    "ABAABBB",
			want: 6,
			k:    2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

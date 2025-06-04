package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kDisinctCharacters(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		k    int
		want int
	}{
		{
			str:  "aababbcaacc",
			k:    2,
			want: 6,
		},
		{
			str:  "abcddefg",
			k:    3,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kDisinctCharacters(tt.str, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

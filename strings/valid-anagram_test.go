package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		{
			name: "test case 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "test case 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

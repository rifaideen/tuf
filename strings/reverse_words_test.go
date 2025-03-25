package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test case 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "test case 2",
			s:    "  hello world  ",
			want: "world hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_reverseWordsSimplified(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test case 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "test case 2",
			s:    "  hello world  ",
			want: "world hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWordsSimplified(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

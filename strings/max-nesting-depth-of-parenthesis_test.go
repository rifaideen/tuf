package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepthStack(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "test case 1",
			s:    "(1+(2*3)+((8)/4))+1",
			want: 3,
		},
		{
			name: "test case 2",
			s:    "(1)+((2))+(((3)))",
			want: 3,
		},
		{
			name: "test case 3",
			s:    "()(())((()()))",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepthStack(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_maxDepthCounter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "test case 1",
			s:    "(1+(2*3)+((8)/4))+1",
			want: 3,
		},
		{
			name: "test case 2",
			s:    "(1)+((2))+(((3)))",
			want: 3,
		},
		{
			name: "test case 3",
			s:    "()(())((()()))",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepthCounter(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

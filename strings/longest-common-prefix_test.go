package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		strs []string
		want string
	}{
		{
			name: "test case 1",
			strs: []string{"leets", "leetcode", "leet", "leeds"},
			want: "lee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_longestCommonPrefixBS(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		strs []string
		want string
	}{
		{
			name: "test case 1",
			strs: []string{"leets", "leetcode", "leet", "leeds"},
			want: "lee",
		},
		{
			name: "test case 2",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "test case 3",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefixBS(tt.strs)
			assert.Equal(t, tt.want, got)
		})
	}
}

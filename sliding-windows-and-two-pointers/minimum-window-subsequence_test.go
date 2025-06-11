package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minWindowSubsequence(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s1   string
		s2   string
		want string
	}{
		{
			name: "test case 1",
			s1:   "XAYMBAZBDCE",
			s2:   "ABE",
			want: "AZBDCE",
		},
		{
			name: "test case 2",
			s1:   "ANOOONANMAOONNANKAONNAN",
			s2:   "ANNAN",
			want: "AONNAN",
		},
		{
			name: "test case 2",
			s1:   "ANOOONANMAOONNANKAONNANKARNNAN",
			s2:   "ANNAN",
			want: "AONNAN",
		},
		{
			s1:   "dynamicprogramming",
			s2:   "mm",
			want: "mm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindowSubsequence(tt.s1, tt.s2)
			assert.Equal(t, tt.want, got)
		})
	}
}

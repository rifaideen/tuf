package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		{
			name: "test case 1",
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			name: "test case 2",
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{
			name: "test case 3",
			s:    "paper",
			t:    "title",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIsomorphic(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

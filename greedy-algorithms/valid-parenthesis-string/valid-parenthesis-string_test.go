package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkValidString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		{
			s:    "()",
			want: true,
		},
		{
			s:    "(*)",
			want: true,
		},
		{
			s:    "(*))",
			want: true,
		},
		{
			s:    ")(*))",
			want: false,
		},
		{
			s:    "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkValidString(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

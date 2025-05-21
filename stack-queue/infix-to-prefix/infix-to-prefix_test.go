package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_infix2prefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expression string
		want       string
	}{
		{
			name:       "test case 1",
			expression: "F+D-C*(B+A)",
			want:       "+-*+ABCDF",
		},
		// {
		// 	name:       "test case 1",
		// 	expression: "(A+B)/(C-D)-(E*F)",
		// 	want:       "AB+CD-/EF*-",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := infix2prefix(tt.expression)
			assert.Equal(t, tt.want, got)
		})
	}
}

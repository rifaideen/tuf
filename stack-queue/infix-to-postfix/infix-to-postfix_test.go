package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_infix2postfix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		expr string
		want string
	}{
		{
			name: "test case 1",
			expr: "(A+B)*(C-D)",
			want: "AB+CD-*",
		},
		{
			name: "test case 1",
			expr: "(A+B)/(C-D)-(E*F)",
			want: "AB+CD-/EF*-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := infix2postfix(tt.expr)
			assert.Equal(t, tt.want, got)
		})
	}
}

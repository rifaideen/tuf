package main

import (
	"fmt"
	"testing"
)

func Test_addOperators(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num    string
		target int
		want   []string
	}{
		{
			num:    "123",
			target: 6,
			want: []string{
				"1+2+3",
				"1*2*3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOperators(tt.num, tt.target)
			fmt.Println(got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			name:   "Example 1",
			num:    "123",
			target: 6,
			want: []string{
				"1+2+3",
				"1*2*3",
			},
		},
		{
			name:   "Example 2",
			num:    "232",
			target: 8,
			want: []string{
				"2+3*2",
				"2*3+2",
			},
		},
		{
			name:   "Leading Zeroes 1",
			num:    "105",
			target: 5,
			want: []string{
				"1*0+5",
				"10-5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOperators(tt.num, tt.target)
			assert.Equal(t, tt.want, got)
			// fmt.Println(got)
		})
	}
}

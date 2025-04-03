package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubstr(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		k    int
		want int
	}{
		{
			name: "test case 1",
			str:  "aba",
			k:    2,
			want: 3,
		},
		{
			name: "test case 2",
			str:  "abaaca",
			k:    1,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubstr(tt.str, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

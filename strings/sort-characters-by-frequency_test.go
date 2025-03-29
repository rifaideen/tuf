package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test case 1",
			s:    "ree",
			want: "eer",
		},
		{
			name: "test case 2",
			s:    "cccaaacc",
			want: "cccccaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findOrder(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		words []string
		k     int
		want  string
	}{
		{
			words: []string{
				"baa",
				"abcd",
				"abca",
				"cab",
				"cad",
			},
			want: "bdac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.words)

			assert.Equal(t, tt.want, got)
		})
	}
}

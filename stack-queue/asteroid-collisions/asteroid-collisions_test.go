package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_asteroidCollision(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		asteroids []int
		want      []int
	}{
		// {
		// 	asteroids: []int{5, 10, -5},
		// 	want:      []int{5, 10},
		// },
		{
			asteroids: []int{8, -8},
			want:      []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			assert.Equal(t, tt.want, got)
		})
	}
}

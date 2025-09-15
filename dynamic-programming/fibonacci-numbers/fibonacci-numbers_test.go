package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibonacci(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n        int
		memoized *[]int
		want     int
	}{
		{
			n:        5,
			memoized: &[]int{-1, -1, -1, -1, -1, -1},
			want:     5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacci(tt.n, tt.memoized)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_fibonacci2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			n:    5,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacci2(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_fibonacci3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			n:    5,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacci3(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pageFault(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		capacity int
		pages    []int
		want     int
	}{
		{
			capacity: 4,
			pages:    []int{5, 0, 1, 3, 2, 4, 1, 0, 5},
			want:     8,
		},
		{
			capacity: 4,
			pages:    []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2},
			want:     6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pageFault(tt.capacity, tt.pages)
			assert.Equal(t, tt.want, got)
		})
	}
}

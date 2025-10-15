package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRodCutting(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		price []int
		n     int
		want  int
	}{
		{
			price: []int{1, 6, 8, 9, 10, 19, 7, 20},
			n:     8,
			want:  25,
		},
		{
			price: []int{1, 5, 8, 9},
			n:     4,
			want:  10,
		},
		{
			price: []int{5, 5, 8, 9, 10, 17, 17, 20},
			n:     8,
			want:  40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RodCutting(tt.price, tt.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRodCuttingMemoized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		price []int
		n     int
		want  int
	}{
		{
			price: []int{1, 6, 8, 9, 10, 19, 7, 20},
			n:     8,
			want:  25,
		},
		{
			price: []int{1, 5, 8, 9},
			n:     4,
			want:  10,
		},
		{
			price: []int{5, 5, 8, 9, 10, 17, 17, 20},
			n:     8,
			want:  40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RodCuttingMemoized(tt.price, tt.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRodCuttingTabulation(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		price []int
		n     int
		want  int
	}{
		{
			price: []int{1, 6, 8, 9, 10, 19, 7, 20},
			n:     8,
			want:  25,
		},
		{
			price: []int{1, 5, 8, 9},
			n:     4,
			want:  10,
		},
		{
			price: []int{5, 5, 8, 9, 10, 17, 17, 20},
			n:     8,
			want:  40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RodCuttingTabulation(tt.price, tt.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRodCuttingOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		price []int
		n     int
		want  int
	}{
		{
			price: []int{1, 6, 8, 9, 10, 19, 7, 20},
			n:     8,
			want:  25,
		},
		{
			price: []int{1, 5, 8, 9},
			n:     4,
			want:  10,
		},
		{
			price: []int{5, 5, 8, 9, 10, 17, 17, 20},
			n:     8,
			want:  40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RodCuttingOptimized(tt.price, tt.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

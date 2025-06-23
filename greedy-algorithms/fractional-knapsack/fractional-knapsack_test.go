package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fractionalKnapsack(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values   []int
		weights  []int
		capacity int
		want     float64
	}{
		// {
		// 	values:   []int{60, 100, 120},
		// 	weights:  []int{10, 20, 30},
		// 	capacity: 50,
		// 	want:     240.0,
		// },
		// {
		// 	values:   []int{10, 20, 30},
		// 	weights:  []int{5, 10, 15},
		// 	capacity: 100,
		// 	want:     60,
		// },
		// {
		// 	values:   []int{124, 456, 789},
		// 	weights:  []int{50, 15, 30},
		// 	capacity: 60,
		// 	want:     1282.2,
		// },
		{
			values:   []int{1, 5, 7, 2, 7, 10},
			weights:  []int{4, 9, 6, 3, 7, 3},
			capacity: 24,
			want:     28.777778,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fractionalKnapsack(tt.values, tt.weights, tt.capacity)
			assert.Equal(t, tt.want, got)
		})
	}
}

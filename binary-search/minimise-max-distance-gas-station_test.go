package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimiseMaxDistance(t *testing.T) {
	distance := MinimiseMaxDistance([]int{1, 2, 3, 4, 5}, 4)
	assert.Equal(t, 0.5, distance)
}

func TestMinimiseMaxDistancePQ(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{
			name: "Test case 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    4,
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimiseMaxDistancePQ(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

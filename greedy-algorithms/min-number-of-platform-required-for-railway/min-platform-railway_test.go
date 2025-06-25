package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPlatform(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arrivals   []int
		departures []int
		want       int
	}{
		{
			arrivals:   []int{900, 940, 950, 1100, 1500, 1800},
			departures: []int{910, 1200, 1120, 1130, 1900, 2000},
			want:       3,
		},
		{
			arrivals:   []int{900, 1235, 1100},
			departures: []int{1000, 1240, 1200},
			want:       1,
		},
		{
			arrivals:   []int{1000, 935, 1100},
			departures: []int{1200, 1240, 1130},
			want:       3,
		},
		{
			arrivals:   []int{1114, 825, 357, 1415, 54},
			departures: []int{1740, 1110, 2238, 1535, 2323},
			want:       4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPlatform(tt.arrivals, tt.departures)
			assert.Equal(t, tt.want, got)
		})
	}
}

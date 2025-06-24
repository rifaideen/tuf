package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxMeetings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		start []int
		end   []int
		want  int
	}{
		{
			start: []int{1, 3, 0, 5, 8, 5},
			end:   []int{2, 4, 6, 7, 9, 9},
			want:  4,
		},
		{
			start: []int{10, 12, 20},
			end:   []int{20, 25, 30},
			want:  1,
		},
		{
			start: []int{1, 2},
			end:   []int{100, 99},
			want:  1,
		},
		{
			start: []int{5, 8, 9, 7, 2, 6, 4},
			end:   []int{2, 1, 5, 3, 4, 8, 6},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxMeetings(tt.start, tt.end)
			assert.Equal(t, tt.want, got)
		})
	}
}

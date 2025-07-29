package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepresent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		m    int
		list [][]int
		want [][]int
	}{
		{
			n: 5,
			m: 6,
			list: [][]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 4},
				{3, 5},
				{4, 5},
			},
			want: [][]int{
				nil,       // 0
				{2, 3},    //1
				{1, 4},    //2
				{1, 4, 5}, //3
				{2, 3, 5}, //4
				{3, 4},    //5
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Represent(tt.n, tt.m, tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

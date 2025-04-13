package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPairs(t *testing.T) {
	tests := []struct {
		name   string
		head   *Node
		target int
		want   [][]int
	}{
		{
			name:   "test case 1",
			head:   constructDll([]int{1, 2, 3, 4, 9}),
			target: 5,
			want: [][]int{
				{1, 4},
				{2, 3},
			},
		},
		{
			name:   "empty list",
			head:   constructDll([]int{}),
			target: 10,
			want:   [][]int{},
		},
		{
			name:   "single node",
			head:   constructDll([]int{5}),
			target: 10,
			want:   [][]int{},
		},
		{
			name:   "two nodes, sum matches",
			head:   constructDll([]int{2, 8}),
			target: 10,
			want:   [][]int{{2, 8}},
		},
		{
			name:   "two nodes, sum doesn't match",
			head:   constructDll([]int{2, 7}),
			target: 10,
			want:   [][]int{},
		},
		{
			name:   "no pairs found",
			head:   constructDll([]int{1, 2, 6, 7}),
			target: 10,
			want:   [][]int{},
		},
		{
			name:   "duplicate values forming a pair",
			head:   constructDll([]int{2, 4, 4, 6}),
			target: 8,
			want: [][]int{
				{2, 6},
				{4, 4},
			},
		},
		{
			name:   "multiple pairs",
			head:   constructDll([]int{1, 2, 3, 4, 5, 6}),
			target: 7,
			want: [][]int{
				{1, 6},
				{2, 5},
				{3, 4},
			},
		},
		{
			name:   "negative numbers",
			head:   constructDll([]int{-3, -1, 2, 4}),
			target: 1,
			want: [][]int{
				{-3, 4},
				{-1, 2},
			},
		},
		{
			name:   "target smaller than smallest values",
			head:   constructDll([]int{5, 6, 7}),
			target: 2,
			want:   [][]int{},
		},
		{
			name:   "target larger than largest values",
			head:   constructDll([]int{1, 2, 3}),
			target: 10,
			want:   [][]int{},
		},
		{
			name:   "list with duplicates, multiple pairs",
			head:   constructDll([]int{1, 2, 2, 3, 4, 4, 5}),
			target: 6,
			want: [][]int{
				{1, 5},
				{2, 4},
				{2, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := findPairs(tt.head, tt.target)

			for key, pair := range pairs {
				assert.Equal(t, tt.want[key][0], pair[0])
				assert.Equal(t, tt.want[key][1], pair[1])
			}
		})
	}
}

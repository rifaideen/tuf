package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteAllOccurences(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head   []int
		target int
		want   []int
	}{
		{
			name:   "test case 1",
			head:   []int{10, 4, 10, 10, 6, 10},
			target: 10,
			want:   []int{4, 6},
		},
		{
			name:   "test case 2",
			head:   []int{10, 10, 10, 10},
			target: 10,
			want:   []int{},
		},
		{
			name:   "test case 3",
			head:   []int{10, 4, 10, 10, 6, 10},
			target: 4,
			want:   []int{10, 10, 10, 6, 10},
		},
		{
			name:   "target not present",
			head:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "single element list - target present",
			head:   []int{5},
			target: 5,
			want:   []int{},
		},
		{
			name:   "single element list - target not present",
			head:   []int{5},
			target: 10,
			want:   []int{5},
		},
		{
			name:   "multiple occurrences at the beginning",
			head:   []int{5, 5, 1, 2, 3},
			target: 5,
			want:   []int{1, 2, 3},
		},
		{
			name:   "multiple occurrences at the end",
			head:   []int{1, 2, 3, 5, 5},
			target: 5,
			want:   []int{1, 2, 3},
		},
		{
			name:   "multiple occurrences in the middle",
			head:   []int{1, 2, 5, 3, 5, 4},
			target: 5,
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "all elements are the target",
			head:   []int{7, 7, 7, 7, 7},
			target: 7,
			want:   []int{},
		},
		{
			name:   "no occurrences",
			head:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "mixed occurrences",
			head:   []int{1, 5, 2, 5, 3, 5, 4},
			target: 5,
			want:   []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := constructDll(tt.head)

			got := deleteAllOccurences(head, tt.target)

			i := 0
			for got != nil {
				assert.Equal(t, tt.want[i], got.Data)
				got = got.Next
				i++
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "test case 1",
			head: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "test case 2",
			head: []int{1, 2, 6},
			want: []int{1, 2, 6},
		},
		{
			name: "test case 3",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "all duplicates",
			head: []int{2, 2, 2, 2, 2},
			want: []int{2},
		},
		{
			name: "duplicates at the beginning",
			head: []int{0, 0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "duplicates at the end",
			head: []int{1, 2, 3, 4, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "interleaved duplicates",
			head: []int{1, 2, 2, 3, 3, 3, 4, 5, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "multiple consecutive duplicates",
			head: []int{1, 1, 1, 5, 5, 8, 8, 8, 8},
			want: []int{1, 5, 8},
		},
		{
			name: "negative numbers with duplicates",
			head: []int{-3, -3, -2, -1, -1, 0, 1},
			want: []int{-3, -2, -1, 0, 1},
		},
		{
			name: "zero and positive duplicates",
			head: []int{0, 0, 1, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			name: "only one unique element",
			head: []int{7, 7, 7, 7},
			want: []int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := constructDll(tt.head)
			got := removeDuplicates(head)

			i := 0

			for got != nil {
				assert.Equal(t, tt.want[i], got.Data)
				got = got.Next
				i++
			}
		})
	}
}

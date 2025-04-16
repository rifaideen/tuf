package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addOne(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head []int
		want []int
	}{
		{
			name: "test case 1",
			head: []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			name: "test case 2",
			head: []int{9, 9, 9},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "test case 3",
			head: []int{0},
			want: []int{1},
		},
		{
			name: "test case 4",
			head: []int{1},
			want: []int{2},
		},
		{
			name: "test case 5",
			head: []int{9},
			want: []int{1, 0},
		},
		{
			name: "test case 6",
			head: []int{1, 9},
			want: []int{2, 0},
		},
		{
			name: "test case 7",
			head: []int{9, 8},
			want: []int{9, 9},
		},
		{
			name: "test case 8",
			head: []int{9, 9},
			want: []int{1, 0, 0},
		},
		{
			name: "test case 9",
			head: []int{1, 0, 0},
			want: []int{1, 0, 1},
		},
		{
			name: "test case 10",
			head: []int{9, 0, 0},
			want: []int{9, 0, 1},
		},
		{
			name: "test case 11",
			head: []int{1, 2, 9},
			want: []int{1, 3, 0},
		},
		{
			name: "test case 12",
			head: []int{9, 9, 1},
			want: []int{9, 9, 2},
		},
		{
			name: "test case 13",
			head: []int{1, 9, 9},
			want: []int{2, 0, 0},
		},
		{
			name: "test case 14",
			head: []int{9, 9, 9, 9},
			want: []int{1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := constructLL(tt.head)

			got := addOne(head)
			i := 0

			for got != nil {
				assert.Equal(t, tt.want[i], got.data)
				got = got.next
				i++
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructDll(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want *Node
	}{
		{
			name: "test case 1",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test case 2",
			nums: []int{1, 0, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := constructDll(tt.nums)

			i := 0
			for head != nil {
				assert.Equal(t, tt.nums[i], head.Data)
				head = head.Next
				i++
			}
		})
	}
}

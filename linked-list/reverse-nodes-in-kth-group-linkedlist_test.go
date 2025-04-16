package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head []int
		k    int
		want []int
	}{
		{
			name: "test case 1",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := constructLL(tt.head)
			got := reverseKGroup(head, tt.k)

			for _, v := range tt.want {
				assert.Equal(t, v, got.data)
				got = got.next
			}
		})
	}
}

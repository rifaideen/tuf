package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flattenList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *LinkedList
		want []int
	}{
		{
			name: "test case 1",
			head: &LinkedList{
				Val: 3,
				Next: &LinkedList{
					Val: 2,
					Next: &LinkedList{
						Val: 1,
						Next: &LinkedList{
							Val: 4,
							Next: &LinkedList{
								Val: 5,
								Child: &LinkedList{
									Val: 6,
									Child: &LinkedList{
										Val: 8,
									},
								},
							},
							Child: &LinkedList{
								Val: 9,
							},
						},
						Child: &LinkedList{
							Val: 7,
							Child: &LinkedList{
								Val: 11,
								Child: &LinkedList{
									Val: 12,
								},
							},
						},
					},
					Child: &LinkedList{
						Val: 10,
					},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flattenList(tt.head)

			for _, v := range tt.want {
				assert.Equal(t, v, got.Val)
				assert.Nil(t, got.Next)
				got = got.Child
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCapacity(t *testing.T) {
	// weights, days, expected
	testcases := [][]any{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15,
		},
		{
			[]int{3, 2, 2, 4, 1, 4}, 3, 6,
		},
		{
			[]int{1, 2, 3, 1, 1}, 4, 3,
		},
	}

	for _, testcase := range testcases {
		output := ShipWithinDays(testcase[0].([]int), testcase[1].(int))
		assert.Equal(t, testcase[2].(int), output)
	}
}

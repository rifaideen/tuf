package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotatedCount(t *testing.T) {
	type testcase struct {
		input  []int
		output int
	}

	testcases := []testcase{
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2, 3},
			output: 4,
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			output: 3,
		},
	}

	for _, testcase := range testcases {
		output := FindRotatedCount(testcase.input)
		assert.Equal(t, testcase.output, output)
	}
}

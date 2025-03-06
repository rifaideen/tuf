package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthRoot(t *testing.T) {
	// n,m,expected
	testcases := [][]int{
		{3, 27, 3},
		{4, 69, -1},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase[2], NthRoot(testcase[0], testcase[1]))
	}
}

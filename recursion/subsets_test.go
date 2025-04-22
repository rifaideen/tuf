package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyArray(t *testing.T) {
	input := []int{}
	expected := [][]int{
		{},
	}
	assert.Equal(t, expected, subsets(input))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitArray(t *testing.T) {
	assert.Equal(t, 18, SplitArray([]int{7, 2, 5, 10, 8}, 2))
	assert.Equal(t, 9, SplitArray([]int{1, 2, 3, 4, 5}, 2))
}

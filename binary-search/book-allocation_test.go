package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPages(t *testing.T) {
	assert.Equal(t, 113, FindPages([]int{12, 34, 67, 90}, 2))
	assert.Equal(t, -1, FindPages([]int{12, 34, 67, 90}, 5))
}

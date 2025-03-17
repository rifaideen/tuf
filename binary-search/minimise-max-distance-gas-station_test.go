package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimiseMaxDistance(t *testing.T) {
	distance := MinimiseMaxDistance([]int{1, 2, 3, 4, 5}, 4)
	assert.Equal(t, 0.5, distance)
}

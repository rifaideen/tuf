package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQRT(t *testing.T) {
	val := FindSqrt(48)
	assert.Equal(t, 6, val)
}

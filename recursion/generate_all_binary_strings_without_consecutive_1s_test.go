package main

import (
	"fmt"
	"testing"
)

func Test_generateBinaryStrings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		k int
	}{
		{
			k: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateBinaryStrings(tt.k)
		})
	}
}

func Test_generateBinaryStringsOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		k int
	}{
		{
			k: 3,
		},
		{
			k: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateBinaryStringsOptimized(tt.k)
			fmt.Println()
		})
	}
}

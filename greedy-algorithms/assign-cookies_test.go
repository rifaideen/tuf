package main

import "testing"

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		g    []int
		s    []int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findContentChildren(tt.g, tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

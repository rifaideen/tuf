package main

import "testing"

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		isConnected [][]int
		want        int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCircleNum(tt.isConnected)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

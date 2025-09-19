package main

import "testing"

func Test_frogJump(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		heights []int
		k       int
		want    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frogJump(tt.heights, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("frogJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_frogJumpMemoized(t *testing.T) {
// 	tests := []struct {
// 		name string // description of this test case
// 		// Named input parameters for target function.
// 		heights []int
// 		k       int
// 		want    int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := frogJumpMemoized(tt.heights, tt.k)
// 			// TODO: update the condition below to compare got with tt.want.
// 			if true {
// 				t.Errorf("frogJumpMemoized() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

package main

import "testing"

func Test_floodFill(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := floodFill(tt.image, tt.sr, tt.sc, tt.color)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

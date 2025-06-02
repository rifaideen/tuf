package main

import "testing"

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fruits []int
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalFruit(tt.fruits)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
